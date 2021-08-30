package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan chan int)
	defer close(queue)

	go func() { // reader

		for {
			ch := <-queue
			for i := range ch {
				fmt.Println(i)
			}
			fmt.Println("Done with this channel")

		}
	}()

	go func() { // writer-1
		ch := make(chan int)
		defer close(ch)
		queue <- ch
		ch <- 4
		ch <- 2
	}()

	go func() { // writer-2
		ch := make(chan int)
		defer close(ch)
		queue <- ch
		ch <- 4
		ch <- 20
	}()
	time.Sleep(time.Second)
}
