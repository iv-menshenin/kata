package loginCounter

import (
	"sync"
	"time"
)

type (
	loggedInMap struct {
		counter map[uint64]int
		tz      time.Time
	}
	mapBucketLogger struct {
		mux     sync.Mutex
		buckets []loggedInMap
	}
)

func (m *mapBucketLogger) login(id uint64) {
	tz := time.Now()
	m.mux.Lock()
	m.remap(tz)
	m.buckets[0].counter[id]++
	m.mux.Unlock()
}

func (m *mapBucketLogger) remap(t time.Time) {
	if len(m.buckets[0].counter) == 0 {
		m.buckets[0].tz = t
		return
	}
	if t.Sub(m.buckets[0].tz) > bucketTime {
		copy(m.buckets[1:], m.buckets[:len(m.buckets)-1])
		m.buckets[0] = loggedInMap{
			counter: make(map[uint64]int, bucketSize),
			tz:      t,
		}
	}
}

func (m *mapBucketLogger) count(id uint64) (count int) {
	m.mux.Lock()
	count = m.buckets[0].counter[id]
	cold := m.buckets[1:]
	m.mux.Unlock()
	for _, bucket := range cold {
		count += bucket.counter[id]
	}
	return count
}

func (m *mapBucketLogger) maxLogged() uint64 {
	var counted = make(map[uint64]int)
	m.mux.Lock()
	for id, cnt := range m.buckets[0].counter {
		counted[id] += cnt
	}
	cold := m.buckets[1:]
	m.mux.Unlock()
	for _, bucket := range cold {
		for id, cnt := range bucket.counter {
			counted[id] += cnt
		}
	}
	var (
		max int
		uid uint64
	)
	for id, cnt := range counted {
		if cnt > max {
			max = cnt
			uid = id
		}
	}
	return uid
}

func NewMapBucketLogger() Counter {
	logger := mapBucketLogger{
		buckets: make([]loggedInMap, 0, bucketsCount),
	}
	for i := 0; i < bucketsCount; i++ {
		logger.buckets = append(logger.buckets, loggedInMap{
			counter: make(map[uint64]int, bucketSize),
			tz:      time.Now(),
		})
	}
	return &logger
}
