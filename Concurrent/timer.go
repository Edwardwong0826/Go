package main

import (
	"fmt"
	"time"
)

func main() {
	// below means wait two seconds
	timer1 := time.NewTimer(time.Second * 2)
	t1 := time.Now()
	fmt.Printf("time now :%v\n", t1)

	// this is channel return by timer
	// so that channel will block and wait the time set by the timer
	t2 := <-timer1.C
	fmt.Printf("time now :%v\n", t2)

	fmt.Printf("time now :%v\n", time.Now())

	timer2 := time.NewTimer(time.Second * 2)
	<-timer2.C

	fmt.Printf("time now :%v\n", time.Now())

	fmt.Printf("time now :%v\n", time.Now())

	// time.After return value is chan Time
	<-time.After(time.Second * 2)

	fmt.Printf("time now :%v\n", time.Now())

	timer3 := time.NewTimer(time.Second)

	go func() {
		<-timer3.C
		fmt.Println("Timer 3 expited")
	}()

	stop := timer3.Stop() // stop the timer so that timer won't block

	if stop {
		fmt.Println("Timer 3 Stopped")
	}

	fmt.Printf("before :%v\n", time.Now())
	timer4 := time.NewTimer(time.Second * 3) // timer originally wait 3 seconds
	timer4.Reset(time.Second * 1)            // reset the timer time
	<-timer4.C
	fmt.Printf("aftter :%v\n", time.Now())

}
