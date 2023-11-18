package main

import (
	"fmt"
	"sync/atomic"
)

func testAddSub() {
	var i int32 = 100
	atomic.AddInt32(&i, 1)
	fmt.Printf("i : %v\n", i)
	atomic.AddInt32(&i, -1)
	fmt.Printf("i : %v\n", i)

	var j int64 = 200
	atomic.AddInt64(&j, 1)
	fmt.Printf("j : %v\n", j)
	atomic.AddInt64(&j, -1)
	fmt.Printf("j : %v\n", j)

}

func testLoadStore() {
	// atomic loadInt32 prevent other goroutines to modify the variable values
	var i int32 = 100
	k := atomic.LoadInt32(&i)
	fmt.Printf("i : %v\n", k)

	atomic.StoreInt32(&i, 200)
	fmt.Printf("i : %v\n", k)
}

// atomic AddInt32 use CAS compare and sawp to compare the value to update
// atomic operation make sure only one goroutine can operates on variable, use atomic can prevent program having many lock operations
// atomic operation
// add
// read
// cas - cas compare variable old value before swap tp new value
// swap - just sawp without comapre variable old value
// write
func main() {

	// cas
	var i int32 = 100
	var p *int32
	p = &i
	b := atomic.CompareAndSwapInt32(&i, *p, 200)
	fmt.Printf("b : %v\n", b)
	fmt.Printf("i : %v\n", i)

}
