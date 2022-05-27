package main

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func Test_goURLS(t *testing.T) {
	type args struct {
		urls    []string
		timeout time.Duration
	}
	tests := []struct {
		name    string
		mock    testDoer
		args    args
		wantErr bool
	}{
		{
			name: "simple",
			mock: testDoer{data: map[string]int{
				"https://foo-bar.com/test1": http.StatusOK,
			}},
			args: args{
				urls:    []string{"https://foo-bar.com/test1"},
				timeout: time.Millisecond * 50,
			},
			wantErr: false,
		},
		{
			name: "two_urls",
			mock: testDoer{data: map[string]int{
				"https://foo-bar.com/test1": http.StatusOK,
				"https://foo-bar.com/test2": http.StatusNotFound,
			}},
			args: args{
				urls:    []string{"https://foo-bar.com/test1", "https://foo-bar.com/test2"},
				timeout: time.Millisecond * 50,
			},
			wantErr: false,
		},
		{
			name: "unknown_url",
			mock: testDoer{data: map[string]int{
				"https://foo-bar.com/test1": http.StatusOK,
			}},
			args: args{
				urls:    []string{"https://foo-bar.com/test1", "https://foo-bar.com/test3"},
				timeout: time.Millisecond * 50,
			},
			wantErr: true,
		},
		{
			name: "empty",
			mock: testDoer{data: map[string]int{}},
			args: args{
				urls:    []string{},
				timeout: time.Millisecond * 50,
			},
			wantErr: false,
		},
		{
			name: "once_timeout",
			mock: testDoer{data: map[string]int{
				"https://foo-bar.com/test1": http.StatusOK,
				"https://foo-bar.com/test2": http.StatusOK,
				"https://foo-bar.com/test3": http.StatusOK,
				"https://foo-bar.com/test4": http.StatusOK,
				"https://foo-bar.com/test5": http.StatusOK,
				"https://foo-bar.com/test6": http.StatusOK,
			}},
			args: args{
				urls: []string{
					"https://foo-bar.com/test1",
					"https://foo-bar.com/test2",
					"https://foo-bar.com/test3",
					"https://foo-bar.com/test4",
					"https://foo-bar.com/test5",
					"https://foo-bar.com/test6",
				},
				timeout: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotErr bool
			var testChecker = tt.mock
			ch := GoURLs(&testChecker, tt.args.urls, tt.args.timeout)
			for resp := range ch {
				if resp.err != nil && !tt.wantErr {
					t.Errorf("got unexpected error: %s", resp.err)
					continue
				}
				if resp.err != nil && tt.wantErr {
					gotErr = true
					break // we break channel reading and test goroutine-leaks here
				}
				if err := testChecker.check(resp.URL, resp.Code); err != nil {
					t.Error(err)
				}
			}
			if tt.wantErr && !gotErr {
				t.Error("expect error")
			}
			if !tt.wantErr && len(testChecker.data) > 0 {
				t.Error("some URLs is missed")
			}
		})
	}
}

type testDoer struct {
	data map[string]int
}

func (t *testDoer) Do(req *http.Request) (*http.Response, error) {
	// simulate net-call
	time.Sleep(time.Millisecond * 5)
	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	default:
		// ok
	}
	if code, ok := t.data[req.URL.String()]; ok {
		return &http.Response{
			Status:     fmt.Sprintf("%d", code),
			StatusCode: code,
		}, nil
	}
	return nil, errors.New("something went wrong")
}

func (t *testDoer) check(URL string, gotCode int) error {
	if respCode, ok := t.data[URL]; ok {
		if respCode == gotCode {
			delete(t.data, URL)
			return nil
		}
		return fmt.Errorf("want: %d, got: %d (URL: %s)", gotCode, respCode, URL)
	}
	return fmt.Errorf("unknown (or requested twice) URL: %s", URL)
}
