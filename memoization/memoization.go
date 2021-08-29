package main

import (
	"fmt"
	"time"
)

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
	memo := &Memo{requests: make(chan request)}
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
	timeout := time.After(exp)
F:
	for {
		var e *entry
		select {
		case <-timeout:
			break F
		case req := <-memo.requests:
			if e == nil {
				e = &entry{ready: make(chan struct{})}
				go e.call(f, req.args)
			}
			go e.deliver(req.response)
		}
	}
	fmt.Println("Exit cache go routine")
	// run a new goroutine that expire the cache
	go memo.serve(f, exp)
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
