package main

import (
	"fmt"
	"math/rand"
	"time"
)

// create int type channel
var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("Send: %v\n", value)
	time.Sleep(time.Second * 3)
	// send value to channel
	values <- value
}

// channel are type thread safe queue, it allow different goroutines to communicate with each others
// <- is called channel operator
// sending and receiving on the channel are blocking operations
// ch <- 69 if we send the value into channel and there is no others goroutine on the other side able to read it out,the code will stop and wait on goroutine until there is reader
// v :=  <- ch on receiving value from channel, if there no goroutine send value on channel, then this part of code will wait until some goroutine to send the value
func main() {

	// get value from channel
	// defer means delay execution until function finish other execution first
	defer close(values)
	go send()
	fmt.Println("Wait...")
	value := <-values
	fmt.Printf("Receive: %v\n", value)
	fmt.Println("End...")

}
