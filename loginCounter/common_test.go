package loginCounter

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func benchCounter(b *testing.B, c Counter) {
	rand.Seed(time.Now().UnixNano())
	b.Helper()
	var (
		wg sync.WaitGroup
		ch = make(chan uint64)
	)
	const (
		threadsCount   = 100
		maxLoggedCount = 5
	)
	var (
		data  = make(map[uint64]int)
		match = make(map[uint64]int)
	)
	for i := 0; i < b.N; {
		cnt := rand.Intn(maxLoggedCount) + 1
		if cnt > b.N-i {
			cnt = b.N - i
		}
		data[rand.Uint64()] += cnt
		i += cnt
	}
	for k, v := range data {
		match[k] = v
	}
	wg.Add(threadsCount)
	for i := 0; i < threadsCount; i++ {
		go func() {
			defer wg.Done()
			var cnt int
			for uid := range ch {
				c.login(uid)
				if cnt++; cnt%777 == 0 {
					c.count(uid)
				}
				if cnt++; cnt%145 == 0 {
					_ = c.maxLogged()
				}
			}
		}()
	}
	b.ResetTimer()
	for {
		if len(data) == 0 {
			break
		}
		var uid uint64
		for uid = range data {
			break
		}
		ch <- uid
		if data[uid]--; data[uid] < 1 {
			delete(data, uid)
		}
	}
	close(ch)
	wg.Wait()
	for uid, v := range match {
		if cnt := c.count(uid); cnt != v {
			b.Errorf("matching data error, expected: %d, got: %d", v, cnt)
		}
	}
}

func testCounterPerformance(t *testing.T, counter Counter) {
	rand.Seed(time.Now().UnixNano())
	const threadsCount = 10000
	const perfBound = 800000
	const testTime = 30
	var qps int64
	var start = time.Now()
	wait := func() {
		for {
			if float64(atomic.LoadInt64(&qps))/time.Since(start).Seconds() < perfBound {
				break
			}
			time.Sleep(time.Microsecond * 5)
		}
		atomic.AddInt64(&qps, 1)
	}
	var iterCounter int32
	var wg sync.WaitGroup
	wg.Add(threadsCount)
	for i := 0; i < threadsCount; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				counter.login(rand.Uint64())
				wait()
				atomic.AddInt32(&iterCounter, 1)
				if time.Since(start).Seconds() > testTime {
					break
				}
			}
		}(&wg)
	}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			counter.count(rand.Uint64())
			counter.maxLogged()
			if time.Since(start).Seconds() > testTime {
				break
			}
			time.Sleep(time.Millisecond * 50)
		}
	}(&wg)
	wg.Wait()
	fmt.Printf("QPS: %d\n", iterCounter/testTime)
	if iterCounter/testTime < 100000 {
		t.Error("low QPS")
	}
	counter.maxLogged()
}
