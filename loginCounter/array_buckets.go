package loginCounter

import (
	"sync"
	"time"
)

type (
	loggedIn struct {
		uid uint64
		tz  time.Time
	}
	arrayBucketLogger struct {
		mux     sync.Mutex
		buckets [][]loggedIn
	}
)

func (m *arrayBucketLogger) login(id uint64) {
	t := time.Now()
	m.mux.Lock()
	m.remap(t)
	m.buckets[0] = append(m.buckets[0], loggedIn{
		uid: id,
		tz:  t,
	})
	m.mux.Unlock()
}

func (m *arrayBucketLogger) remap(t time.Time) {
	if len(m.buckets[0]) == 0 {
		return
	}
	if t.Sub(m.buckets[0][0].tz) > bucketTime {
		var b = m.buckets[len(m.buckets)-1]
		copy(m.buckets[1:], m.buckets[:len(m.buckets)-1])
		m.buckets[0] = b[:0]
	}
}

func (m *arrayBucketLogger) count(id uint64) (count int) {
	m.mux.Lock()
	var buckets = m.buckets[1 : len(m.buckets)-2]
	for _, el := range m.buckets[0] {
		if el.uid == id {
			count++
		}
	}
	for _, el := range m.buckets[len(m.buckets)-1] {
		if el.uid == id && time.Since(el.tz) < actualTime {
			count++
		}
	}
	m.mux.Unlock()
	for _, bucket := range buckets {
		for _, el := range bucket {
			if el.uid == id {
				count++
			}
		}
	}
	return count
}

func (m *arrayBucketLogger) maxLogged() uint64 {
	var id uint64
	var loginsCounter = make(map[uint64]int)
	m.mux.Lock()
	for _, bucket := range m.buckets {
		for _, el := range bucket {
			loginsCounter[el.uid]++
		}
	}
	m.mux.Unlock()
	var maxLogins = -1
	for uid, loginsCount := range loginsCounter {
		if maxLogins < loginsCount {
			id = uid
			maxLogins = loginsCount
		}
	}
	return id
}

func NewArrayBucketLogger() Counter {
	logger := arrayBucketLogger{
		buckets: make([][]loggedIn, 0, bucketsCount),
	}
	for i := 0; i < bucketsCount; i++ {
		logger.buckets = append(logger.buckets, make([]loggedIn, 0, bucketSize))
	}
	return &logger
}
