package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Goroutine
// when main thread run finished, it doesn't care other child coroutines execution and will exit
func main() {

	// get computer cpu amount use by go
	cpuNumber := runtime.NumCPU()
	fmt.Println("CPU Number : ", cpuNumber)

	// set how cpu to use in go
	runtime.GOMAXPROCS(cpuNumber - 1)

	// go keyworad for parallel run for main thread and child coroutines
	wg.Add(1) // coroutine counter +1
	go test1()
	wg.Add(1)
	go test2()

	wg.Wait() // wait coroutines execution finished
	fmt.Println("main thread exit")

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go test3(i)
	}
	wg.Wait()

}

func test1() {

	for i := 0; i < 10; i++ {
		fmt.Println("test 1() - Hello Go", i)
		time.Sleep(time.Millisecond * 10)
	}
	wg.Done() // coroutine counter -1
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test 2() - Hello Go", i)
		time.Sleep(time.Millisecond * 20)
	}
	wg.Done() // coroutine counter -1
}

func test3(num int) {

	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("coroutine %v print %v data\n", num, i)
		time.Sleep(time.Millisecond * 50)
	}

}
