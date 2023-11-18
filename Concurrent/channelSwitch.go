package main

import (
	"fmt"
	"time"
)

// create int type channel
var chanInt = make(chan int, 0)

var chanStr = make(chan string, 0)

func main() {

	go func() {
		chanInt <- 100
		chanStr <- "hello"
		// without below defer close, once channel read all the values, it will keep fall into default
		defer close(chanInt)
		defer close(chanStr)
	}()

	// channel oen of the example is  多路复用 / 异步io
	// select witch case must a channel opration
	for {
		select {
		case r := <-chanInt:
			fmt.Printf("chanInt: %v\n", r)
		case r := <-chanStr:
			fmt.Printf("chanStr: %v\n", r)
		// without default, after channel swtich read all of the value, select will block until same case channel can be run
		default:
			fmt.Println("default...")
		}

		time.Sleep(time.Second)
	}
}
