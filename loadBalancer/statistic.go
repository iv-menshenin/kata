package loadBalancer

import (
	"sync"
	"sync/atomic"
	"time"
)

type info struct {
	once    sync.Once
	mux     sync.RWMutex
	buckets [5]bucket
}

type bucket struct {
	err int64
	cnt int64
	exe int64
	lat int64
}

func (i *info) setErr() {
	i.mux.RLock()
	atomic.AddInt64(&i.buckets[0].err, 1)
	i.mux.RUnlock()
}

func (i *info) setCnt() {
	i.once.Do(func() {
		go func() {
			for {
				<-time.After(time.Second)
				i.mux.Lock()
				copy(i.buckets[1:], i.buckets[:len(i.buckets)-1])
				i.buckets[0] = bucket{}
				i.mux.Unlock()
			}
		}()
	})
	i.mux.RLock()
	atomic.AddInt64(&i.buckets[0].cnt, 1)
	i.mux.RUnlock()
}

func (i *info) registerQs(d int64) {
	i.mux.RLock()
	atomic.AddInt64(&i.buckets[0].lat, d)
	i.mux.RUnlock()
}

func (i *info) done() {
	i.mux.RLock()
	atomic.AddInt64(&i.buckets[0].exe, 1)
	i.mux.RUnlock()
}

func (i *info) stat() stat {
	i.mux.RLock()
	var s stat
	var ex, la int64
	for _, b := range i.buckets {
		ex += b.exe
		la += b.lat
		s.cnt += b.cnt
		s.err += b.err
	}
	i.mux.RUnlock()
	if s.cnt > ex {
		s.lat = la / s.cnt
	} else if ex > 0 {
		s.lat = la / ex
	}
	s.exe = ex
	return s
}
