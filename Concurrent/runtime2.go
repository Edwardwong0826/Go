package main

import (
	"fmt"
	"runtime"
	"time"
)

func show(s string) {

	for i := 0; i < 10; i++ {
		fmt.Println(s)
		if i >= 5 {
			// like loop break, it will exit current coroutine
			runtime.Goexit()
		}
	}

}

func main() {

	go show("Java")
	time.Sleep(time.Second)

}
