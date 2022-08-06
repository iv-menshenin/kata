package loadBalancer

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func Test_Balancer(t *testing.T) {
	const testThreads = 100
	var (
		b          = New(getDoers(), &roundRobin{})
		wg         sync.WaitGroup
		errCounter int64
		allCounter int64
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	wg.Add(testThreads)
	for i := 0; i < testThreads; i++ {
		go func() {
			<-time.After(time.Millisecond * time.Duration(20))
			defer wg.Done()
			for {
				atomic.AddInt64(&allCounter, 1)
				req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/test", nil)
				if resp, err := b.Do(req); err != nil || resp.StatusCode > 399 {
					atomic.AddInt64(&errCounter, 1)
				}
				select {
				case <-ctx.Done():
					return

				case <-time.After(time.Millisecond * 20):
					// another one
				}
			}
		}()
	}

	<-ctx.Done()
	wg.Wait()
	fmt.Printf("done %d requests\nerrors: %d\nrate: %0.4f\n", allCounter, errCounter, float64(allCounter-errCounter)/float64(allCounter))
	fmt.Printf("%v\n", chosen[:len(b.doers)])
	for i := range b.info[:len(b.doers)] {
		fmt.Printf("%+v\n", b.info[i].buckets)
	}
}

func getDoers() []httpDoer {
	return []httpDoer{
		&goodDoer{},
		&goodDoer{},
		&longStarterDoer{},
		&longStarterDoer{},
		&surpriseDoer{},
		&surpriseDoer{},
		&badDoer{},
		&badDoer{},
		&badServer{},
		&badServer{},
	}
}

/*
========================
BEFORE longStarted
 --- ideal
done 74476 requests
errors: 0
rate: 1.0000

 --- best of roundRobin
done 955 requests
errors: 122
rate: 0.8723

-- best of judge
done 37118 requests
errors: 129
rate: 0.9965

========================
AFTER longStarted
 --- best of roundRobin
done 11728 requests
errors: 5533
rate: 0.5282

-- best of judge
done 21283 requests
errors: 488
rate: 0.9771

*/
