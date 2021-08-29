package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {
	memo := NewMemoize(sometestfunction, 1*time.Second)
	timeout := time.After(10 * time.Second)
	for i := 0; i < 5; i++ {
		go func(i int) {
			data, _ := memo.Get(fmt.Sprintf("some data"))
			fmt.Printf("Data %v %v \n", data, i)
		}(i)
	}
	time.Sleep(2 * time.Second)
	for i := 5; i < 10; i++ {
		go func(y int) {
			data, _ := memo.Get(fmt.Sprintf("some more data"))
			fmt.Printf("Data %v %v \n", data, y)
		}(i)
	}
	time.Sleep(2 * time.Second)
	for i := 10; i < 15; i++ {
		go func(i int) {
			data, _ := memo.Get(fmt.Sprintf("another type of data"))
			fmt.Printf("Data %v %v \n", data, i)
		}(i)
	}

	<-timeout
}

type usageDetails struct {
	Usage struct {
		SrvUid  string `json:"server-uid"`
		Enabled string `json:"enabled"`
		License string `json:"license-id"`
		Crn     string `json:"crn"`
	} `json:"usage"`
}

func sometestfunction(testkey ...string) (interface{}, error) {
	return testkey, nil
}

func getStatsFromServer(authKey ...string) (interface{}, error) {
	client := http.Client{Timeout: 30 * time.Second}
	url, _ := url.Parse(fmt.Sprintf("http://localhost:9191/api/usage?Authorization=%s", url.QueryEscape(authKey[0])))

	req := &http.Request{URL: url, Method: "GET"}
	res, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Failed to get data from the MF server, %v", err)
		return nil, errors.New(fmt.Sprintf("failed to get data from MF server: %v", err))
	}

	if res.StatusCode != 200 {
		fmt.Errorf("Failed to read data from the MF Health API.")
		return nil, errors.New(fmt.Sprintf("failed to get data from MF server, %v", err))
	}

	data := &usageDetails{}
	json.NewDecoder(res.Body).Decode(data)
	return data, nil
}

// memoization.go
type result struct {
	value interface{}
	err   error
}

type request struct {
	args     []string
	response chan result
}

type entry struct {
	result  *result
	ready   chan struct{}
	refresh chan struct{}
}

// function we want to memoize and will fetch the data and put the result
type Func func(args ...string) (interface{}, error)

type Memo struct {
	requests chan request
}

/**
* f 	- the function we wish to memoize
* exp 	- the time to expire the cache
* returns a struct that contains the channel to receive incoming requests
 */
func NewMemoize(f Func, exp time.Duration) *Memo {
	memo := &Memo{requests: make(chan request, 1)}
	go memo.serve(f, exp)
	return memo
}

// args can be any number of array strings
func (memo *Memo) Get(args ...string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{args, response}
	res := <-response // wait for the result and than pass it
	return res.value, res.err
}

func (memo *Memo) serve(f Func, exp time.Duration) {
	// timeout := time.After(exp)
	ticker := time.NewTicker(exp)
	var e *entry
	// F:
	for {
		select {
		// case <-timeout:
		// break F
		case <-ticker.C:
			if e == nil {
				return
			}
			select {
			case <-e.ready:
				e.ready = make(chan struct{})
			case req := <-memo.requests:
				go e.call(f, req.args)
				go e.deliver(req.response)
			}
		case req := <-memo.requests:
			if e == nil {
				fmt.Println("Call to get entry")
				e = &entry{ready: make(chan struct{})}
				go e.call(f, req.args)
			}
			go e.deliver(req.response)
		}
	}
	// fmt.Println("Exit cache go routine")
	// run a new goroutine that expire the cache
	// go memo.serve(f, exp)
}

func (e *entry) call(f Func, args []string) {
	// Evaluate the function
	e.result = &result{}
	e.result.value, e.result.err = f(args...)
	// Broadcast the ready condition
	close(e.ready)
	// <-e.refresh // wait till e refreshes
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition, if the channel is closed, then all the goroutines will skip this
	<-e.ready
	// Send the result to the client
	response <- *e.result
}
