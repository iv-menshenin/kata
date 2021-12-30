package loginCounter

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

const (
	maxTestLoggedPerUserCount = 5
	clickersRoutinesCount     = 100
	performanceTestDataSize   = 1000000
	functionalTestDataSize    = 100000
)

func benchCounterLogin(b *testing.B, c Counter) {
	var wg sync.WaitGroup
	var ch = make(chan uint64)
	var data, _ = makeClickData(b.N)

	startClickers(ch, c, &wg)
	stopAdmin := startAdmin(c, &wg)

	b.ResetTimer()
	sendData(ch, data)
	close(ch)
	stopAdmin()
	wg.Wait()
}

func benchCounterCount(b *testing.B, c Counter) {
	var wg sync.WaitGroup
	var ch = make(chan uint64)
	var data, match = makeClickData(b.N)

	startClickers(ch, c, &wg)
	stopAdmin := startAdmin(c, &wg)

	sendData(ch, data)
	close(ch)
	stopAdmin()
	wg.Wait()

	b.ResetTimer()
	checkCounterResults(b, match, c)
}

func testCounterPerformance(t *testing.T, c Counter) {
	var wg sync.WaitGroup
	var ch = make(chan uint64)
	var data, _ = makeClickData(performanceTestDataSize)

	startClickers(ch, c, &wg)
	stopAdmin := startAdmin(c, &wg)

	started := time.Now()
	sendData(ch, data)
	close(ch)
	stopAdmin()
	wg.Wait()
	doneTime := time.Since(started)
	perf := performanceTestDataSize / doneTime.Seconds()

	fmt.Printf("QPS: %0.3f\n", perf)
	if perf < 100000 {
		t.Error("low QPS")
		return
	}
}

func testCounterCountFunc(t *testing.T, c Counter) {
	var wg sync.WaitGroup
	var ch = make(chan uint64)
	var data, match = makeClickData(functionalTestDataSize)

	startClickers(ch, c, &wg)
	stopAdmin := startAdmin(c, &wg)

	sendData(ch, data)
	close(ch)
	stopAdmin()
	wg.Wait()

	checkCounterResults(t, match, c)
}

func makeClickData(clicksCount int) (data, copy map[uint64]int) {
	rand.Seed(time.Now().UnixNano())
	data = make(map[uint64]int)
	copy = make(map[uint64]int)
	for i := 0; i < clicksCount; {
		loggedCnt := rand.Intn(maxTestLoggedPerUserCount) + 1
		if loggedCnt > clicksCount-i {
			loggedCnt = clicksCount - i
		}
		data[rand.Uint64()] += loggedCnt
		i += loggedCnt
	}
	for k, v := range data {
		copy[k] = v
	}
	return data, copy
}

func clickWorker(uidChan <-chan uint64, c Counter) {
	for uid := range uidChan {
		c.login(uid)
	}
}

func admRoutinesWorker(t *time.Ticker, stopTimer <-chan struct{}, c Counter) {
	for {
		select {
		case <-t.C:
			c.count(rand.Uint64())
			c.maxLogged()
		case <-stopTimer:
			return
		}
	}
}

func startClickers(ch <-chan uint64, c Counter, wg *sync.WaitGroup) {
	wg.Add(clickersRoutinesCount)
	for i := 0; i < clickersRoutinesCount; i++ {
		go func() {
			defer wg.Done()
			clickWorker(ch, c)
		}()
	}
}

func startAdmin(c Counter, wg *sync.WaitGroup) func() {
	wg.Add(1)
	var ticker = time.NewTicker(time.Millisecond * 250)
	var stopTimer = make(chan struct{})
	go func() {
		defer wg.Done()
		admRoutinesWorker(ticker, stopTimer, c)
	}()
	return func() {
		close(stopTimer)
		ticker.Stop()
	}
}

type errorfer interface {
	Errorf(format string, args ...interface{})
}

func checkCounterResults(e errorfer, data map[uint64]int, c Counter) {
	var wg sync.WaitGroup
	var maxCnt int
	var maxUID uint64
	for uid, v := range data {
		if v > maxCnt {
			maxCnt = v
			maxUID = uid
		}
		wg.Add(1)
		go func(uid uint64, v int) {
			defer wg.Done()
			if cnt := c.count(uid); cnt != v {
				e.Errorf("matching data error, expected: %d, got: %d", v, cnt)
			}
		}(uid, v)
	}
	wg.Wait()
	if maxGot := c.maxLogged(); maxCnt != c.count(maxGot) {
		e.Errorf("maxLogged() error, expected: %d, got: %d", maxUID, maxGot)
	}
}

func sendData(ch chan<- uint64, data map[uint64]int) {
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
}
