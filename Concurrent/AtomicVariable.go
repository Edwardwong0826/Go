package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// using lock to synchronize

// var i = 100

// var lock sync.Mutex

// func add() {
// 	lock.Lock()
// 	i++
// 	lock.Unlock()
// }

// func sub() {
// 	lock.Lock()
// 	i--
// 	lock.Unlock()
// }

var i int32 = 100

// atomic AddInt32 use CAS compare and sawp to compare the value to update
// atomic operation make sure only one goroutine can operates on variable, use atomic can prevent program having many lock operations
// atomic operation
// add
// read
// cas
// swap
// write
func add() {
	atomic.AddInt32(&i, 1)
}

func sub() {
	atomic.AddInt32(&i, -1)
}

func main() {

	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}

	time.Sleep(time.Second * 3)
	fmt.Printf("i : %v\n", i)
}
