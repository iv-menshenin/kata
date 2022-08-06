package loadBalancer

import (
	"math/rand"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

type (
	arbitrator interface {
		choose([]stat) int
	}
	stat struct {
		cnt int64
		err int64
		lat int64
		exe int64
	}
)

type roundRobin struct {
	next int64
}

func (r *roundRobin) choose(doers []stat) int {
	var trying = len(doers)
	for {
		trying--
		i := int(atomic.AddInt64(&r.next, 1) % int64(len(doers)))
		if trying > 0 && doers[i].err > 5 {
			continue
		}
		return i
	}
}

type judge struct {
	blocked  map[int]time.Time
	halfOpen map[int]int64
	once     sync.Once
	mx       sync.Mutex
	counter  int64
}

func (g *judge) choose(doers []stat) int {
	g.once.Do(func() {
		g.blocked = make(map[int]time.Time)
		g.halfOpen = make(map[int]int64)
	})
	var buff [16]int
	var clients = g.healthyClients(doers, buff[:])

	sort.Slice(clients, func(i, j int) bool {
		lessLatency := doers[clients[i]].lat < doers[clients[j]].lat
		eqLatency := doers[clients[i]].lat == doers[clients[j]].lat // latency can be zero
		lessCount := doers[clients[i]].cnt < doers[clients[j]].cnt
		return lessLatency || (eqLatency && lessCount)
	})
	if len(clients) > 2 {
		clients = clients[:2]
	}
	clientNum := rand.Intn(len(clients))
	return clients[clientNum]
}

func (g *judge) healthyClients(doers []stat, buff []int) []int {
	c := atomic.AddInt64(&g.counter, 1)
	g.mx.Lock()
	var clients = buff[:0]
	for i, doer := range doers {
		bt, bl := g.blocked[i]
		ht, ho := g.halfOpen[i]
		// have to open circuit if there is more than 5 errors
		if doer.err > 5 {
			if !bl {
				g.blocked[i] = time.Now()
				bl = true
			}
			if ho {
				delete(g.halfOpen, i)
				ho = false
			}
		}
		// have to close circuit after 10 successful requests
		if ho && doer.exe-ht > 10 {
			if doer.err == 0 {
				delete(g.halfOpen, i)
				ho = false
			}
		}
		// move to half-opened after 2 sec
		if !ho && bl && time.Since(bt).Seconds() > 2 {
			g.halfOpen[i] = doer.exe
			delete(g.blocked, i)
			bl, ho = false, true
		}
		if !bl && (!ho || canWeTryHalfOpened(c, i)) {
			clients = append(clients, i)
		}
	}
	// no healthy clients? take those who are not blocked
	if len(clients) == 0 {
		for i := range doers {
			_, bl := g.blocked[i]
			_, ho := g.halfOpen[i]
			if !bl || ho {
				clients = append(clients, i)
			}
		}
	}
	// all blocked? accept all
	if len(clients) == 0 {
		for i := range doers {
			clients = append(clients, i)
		}
	}
	g.mx.Unlock()
	return clients
}

func canWeTryHalfOpened(queryNum int64, clientNum int) bool {
	return int(queryNum%100) == clientNum*10
}
