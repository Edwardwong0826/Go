package main

import (
	"fmt"
	"runtime"
)

func show(s string) {

	for i := 0; i < 2; i++ {
		fmt.Println(s)
	}

}

// runtime package provided by go to manage coroutine related api
func main() {
	go show("Java")

	for i := 0; i < 2; i++ {
		runtime.Gosched() // let other coroutine to execute first
		fmt.Println("Go")
	}

}
