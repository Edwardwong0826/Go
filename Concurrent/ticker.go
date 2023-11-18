package main

import (
	"fmt"
	"time"
)

// Timer only excute once, Ticker can execute interval to stop time
func main() {
	ticker := time.NewTicker(time.Second)
	//counter := 1

	// for _ = range ticker.C {
	// 	fmt.Println("Ticker 1")
	// 	counter++

	// 	if counter >= 5 {
	// 		break
	// 	}
	// }

	// ticker.Stop()
	// this is a unbuffered channel
	// by default, sends and received block until the other side is ready
	// this allows goroutines to synchronize without explicit locks or condition variables
	chanInt := make(chan int)
	// this is buffered channel
	// send to buffered channel block only when the buffer is full
	// received block when the buffer is empty
	// chanInt := make(chan int, 100)
	go func() {

		for _ = range ticker.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}

		}
	}()

	sum := 0
	for v := range chanInt {
		fmt.Printf("Receiving %v\n", v)
		sum += v
		if sum >= 10 {
			fmt.Printf("Sum %v\n", sum)
			break
		}

	}
}
