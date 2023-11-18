package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100

var lock sync.Mutex

var wg sync.WaitGroup

func add() {

	defer wg.Done() // wt.Done() is actually wt.Add(-1)
	lock.Lock()
	i += 1
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}

func sub() {

	defer wg.Done()
	lock.Lock()
	time.Sleep(time.Millisecond * 2)

	i -= 1
	lock.Unlock()
}

func main() {

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()

	}

	wg.Wait()

	fmt.Println(i)
}
