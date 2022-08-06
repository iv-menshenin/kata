package loadBalancer

import (
	"context"
	"net/http"
	"sync/atomic"
	"time"
)

type client struct {
	arbitrator arbitrator
	doers      []httpDoer
	info       []info
	b          chan struct{}
}

func New(doers []httpDoer, a arbitrator) *client {
	c := client{
		arbitrator: a,
		doers:      doers,
		info:       make([]info, len(doers)),
		b:          make(chan struct{}, 100),
	}
	return &c
}

func (c *client) Do(r *http.Request) (*http.Response, error) {
	c.b <- struct{}{}
	defer func() {
		<-c.b
	}()

	ctx, cancel := context.WithTimeout(r.Context(), time.Second)
	defer cancel()

	var done = make(chan struct{})
	var res *http.Response
	var err error
	go func() {
		defer close(done)
		res, err = c.getDoer().Do(r.Clone(ctx))
	}()

	select {

	case <-done:
		return res, err

	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

var chosen [64]int64

func (c *client) getDoer() httpDoer {
	var (
		max  = len(c.doers)
		buck = make([]stat, max)
	)
	for n := 0; n < len(c.doers); n++ {
		buck[n] = c.info[n].stat()
	}
	curr := c.arbitrator.choose(buck)
	atomic.AddInt64(&chosen[curr], 1)

	return wrappedDoer{
		doer: c.doers[curr],
		info: &c.info[curr],
	}
}

type wrappedDoer struct {
	start time.Time
	delta time.Duration
	doer  httpDoer
	info  *info
}

func (d wrappedDoer) Do(r *http.Request) (*http.Response, error) {
	d.info.setCnt()
	defer close(d.monitor())
	resp, err := d.doer.Do(r)
	d.info.done()
	if err != nil || resp.StatusCode > 499 {
		d.info.setErr()
	}
	return resp, err
}

func (d wrappedDoer) monitor() chan<- struct{} {
	d.delta = 0
	d.start = time.Now()
	var done = make(chan struct{})
	go func() {
		for {
			select {

			case <-time.After(time.Millisecond):
				d.tick()

			case <-done:
				d.tick()
				return
			}
		}
	}()
	return done
}

func (d wrappedDoer) tick() {
	latency := time.Since(d.start) - d.delta
	d.info.registerQs(latency.Nanoseconds())
	d.delta += latency
}
