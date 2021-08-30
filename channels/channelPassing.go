package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	queue := make(chan int)
	go printFormat(nil, queue)
	queue <- 1
	queue <- 2
	close(queue)
	// fake some actions in between
	time.Sleep(2 * time.Second)
	queue = make(chan int)
	go printFormat(cancel, queue)
	queue <- 3
	queue <- 4
	close(queue)
	<-ctx.Done()
}

func printFormat(cancel context.CancelFunc, q chan int) {
	for i := range q {
		fmt.Printf("Data %d \n", i)
	}
	if cancel != nil {
		cancel()
	}
}
