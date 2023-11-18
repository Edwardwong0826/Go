package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Goroutine
// when main thread run finished, it doesn't care other child coroutines execution and will exit
// example calculate 1-120000 number which one is prime number ? use goroutine to count
func main() {

	for i := 1; i < 4; i++ {
		wg.Add(1)
		go test(i)
	}

	wg.Wait()
	fmt.Println("main thread exit")

}

func test(n int) {
	for num := (n-1)*30 + 1; num < n*30; num++ {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
			if flag {
				fmt.Println(num, " is prime number")
			}
		}
	}
	wg.Done()
}
