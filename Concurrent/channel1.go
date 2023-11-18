package main

import (
	"fmt"
)

// create int type channel
var c = make(chan int)

func main() {

	go func() {
		for i := 0; i < 2; i++ {
			c <- i
		}
		// so we have to close the channel, when outside read value from channel more than
		// its store value, it will return default value for the channel value type
		close(c)
	}()

	r := <-c
	fmt.Printf("r: %v\n", r)

	r = <-c
	fmt.Printf("r: %v\n", r)

	// when read the value from channel more than its store value, it will throw deadlock
	r = <-c
	fmt.Printf("r: %v\n", r)

	// same as above, just enhance for loop, go call for range
	// for {
	// 	v, ok := <-c
	// 	if ok {
	// 		fmt.Printf("v: %v\n", v)
	// 	} else {
	// 		break
	// 	}
	// }
}
