package loadBalancer

import (
	"errors"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type httpDoer interface {
	Do(r *http.Request) (*http.Response, error)
}

type goodDoer struct {
	parallel int64
}

func (g *goodDoer) Do(r *http.Request) (*http.Response, error) {
	p := atomic.AddInt64(&g.parallel, 1)
	defer atomic.AddInt64(&g.parallel, -1)

	if p > 30 {
		time.Sleep(time.Millisecond * 100)
	}
	if p > 50 {
		time.Sleep(time.Millisecond * 450)
	}
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(80)+20))
	return &http.Response{
		Status:     "200",
		StatusCode: http.StatusOK,
		Proto:      "http",
	}, nil
}

type badDoer struct{}

func (g *badDoer) Do(*http.Request) (*http.Response, error) {
	time.Sleep(time.Second * 30)
	return nil, errors.New("bad connection")
}

type badServer struct{}

func (g *badServer) Do(*http.Request) (*http.Response, error) {
	time.Sleep(time.Millisecond * 60)
	return &http.Response{StatusCode: 500}, nil
}

type surpriseDoer struct {
	cnt  int64
	good goodDoer
}

func (g *surpriseDoer) Do(r *http.Request) (*http.Response, error) {
	if atomic.AddInt64(&g.cnt, 1) > 100 {
		time.Sleep(time.Millisecond * 60)
	}
	if atomic.AddInt64(&g.cnt, 1) > 250 {
		time.Sleep(time.Millisecond * 120)
	}
	if atomic.AddInt64(&g.cnt, 1) > 300 {
		time.Sleep(time.Second * 30)
	}
	return g.good.Do(r)
}

type longStarterDoer struct {
	on   sync.Once
	good goodDoer
}

func (g *longStarterDoer) Do(r *http.Request) (*http.Response, error) {
	g.on.Do(func() {
		time.Sleep(30 * time.Second)
	})
	return g.good.Do(r)
}
