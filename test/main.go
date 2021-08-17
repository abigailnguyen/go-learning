package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main() terminated.")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	// abort := make(chan bool)
	updateTicker := make(chan bool)
	go func() {
		defer fmt.Println("Finished goroutine 1")
		// time.Sleep(5 * time.Second)
		// time.Sleep(10 * time.Second)
		// updateTicker <- true
		close(updateTicker)
	}()

L:
	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("Time is %v\n", t)
		case <-updateTicker:
			fmt.Println("Ticker updated!")
			break L
		}
	}
	updateTicker <- true
}
