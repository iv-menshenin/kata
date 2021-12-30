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
	m.remapIfNeeded(t)
	m.buckets[0] = append(m.buckets[0], loggedIn{
		uid: id,
		tz:  t,
	})
	m.mux.Unlock()
}

func (m *arrayBucketLogger) remapIfNeeded(t time.Time) {
	if len(m.buckets[0]) == 0 {
		return
	}
	if t.Sub(m.buckets[0][0].tz) > bucketTime {
		var b = m.buckets[len(m.buckets)-1]
		copy(m.buckets[1:], m.buckets[:len(m.buckets)-1])
		m.buckets[0] = b[:0]
	}
}

func (m *arrayBucketLogger) getColdBuckets() [][]loggedIn {
	var cold = make([][]loggedIn, len(m.buckets)-2)
	copy(cold[:], m.buckets[1:len(m.buckets)-2])
	return cold
}

func (m *arrayBucketLogger) count(id uint64) (count int) {
	m.mux.Lock()
	// take buckets that could not be changed concurrently
	cold := m.getColdBuckets()
	// let's calculate the data in the hot bucket
	for _, el := range m.buckets[0] {
		if el.uid == id {
			count++
		}
	}
	// this bucket can be reduced in remapIfNeeded
	for _, el := range m.buckets[len(m.buckets)-1] {
		if el.uid == id && time.Since(el.tz) < actualTime {
			count++
		}
	}
	m.mux.Unlock()
	// we can calculate it concurrently without locks
	for _, bucket := range cold {
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
	cold := m.getColdBuckets()
	// let's calculate the data in the hot bucket
	for _, el := range m.buckets[0] {
		loginsCounter[el.uid]++
	}
	// this bucket can be reduced in remapIfNeeded
	for _, el := range m.buckets[len(m.buckets)-1] {
		if el.uid == id && time.Since(el.tz) < actualTime {
			loginsCounter[el.uid]++
		}
	}
	m.mux.Unlock()
	for _, bucket := range cold {
		for _, el := range bucket {
			loginsCounter[el.uid]++
		}
	}
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
