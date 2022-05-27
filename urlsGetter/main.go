package main

import (
	"context"
	"net/http"
	"sync"
	"time"
)

// A list of URLs is given, we need to execute them in parallel with a limit on the number of simultaneous executions.
// Task specifics:
// * Return HTTP response code, and return an error if an error occurs;
// * Set timeout for execution of the whole task;
// * Limit the number of concurrent queries;

type (
	Doer interface {
		Do(req *http.Request) (*http.Response, error)
	}
	queryResult struct {
		err  error
		URL  string
		Code int
	}
)

const limitWorkers = 10

func GoURLs(client Doer, urls []string, timeout time.Duration) <-chan queryResult {
	if client == nil {
		client = http.DefaultClient
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	var ch = make(chan queryResult)
	var wg sync.WaitGroup
	var limiter = make(chan struct{}, limitWorkers)

	wg.Add(len(urls))
	for _, URL := range urls {
		go execTaskAsyncWithLimit(ctx, client, URL, limiter, ch, &wg)
	}

	go func() {
		wg.Wait()
		cancel()
		close(ch)
	}()
	return ch
}

func execTaskAsyncWithLimit(ctx context.Context, client Doer, URL string, limiter chan struct{}, result chan queryResult, wg *sync.WaitGroup) {
	limiter <- struct{}{}
	defer func() {
		<-limiter
		wg.Done()
	}()

	code, err := goURL(ctx, client, URL)
	select {

	case result <- queryResult{err: err, URL: URL, Code: code}:
		return

	case <-ctx.Done():
		return
	}
}

func goURL(ctx context.Context, client Doer, URL string) (int, error) {
	select {

	case <-ctx.Done():
		return 0, ctx.Err()

	default:
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
		if err != nil {
			return 0, err
		}
		resp, err := client.Do(req)
		if err != nil {
			return 0, err
		}
		return resp.StatusCode, nil
	}
}
