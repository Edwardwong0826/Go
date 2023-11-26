package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Hello struct {
}

// Go had integrate its own RPC package, net and net/rpc is the one
// define one remote call function, function normally put into struct inside

// Go RPC rules
// 1. method only can have 2 able serialize parameter means cannot be channel and func type, and second parameter need to be pointer type
// 2. method need to return one error type, and must be public method, means first letter of method need to be capital letter
func (this Hello) SayHello(request string, response *string) error {
	*response = "hello " + request
	return nil
}

func main() {

	fmt.Println("Hello server")
	//1. Register RPC service
	err := rpc.RegisterName("hello", new(Hello))

	if err != nil {
		fmt.Println(err)
	}

	// 2. lister port
	listener, err1 := net.Listen("tcp", "127.0.0.1:8090")

	if err1 != nil {
		fmt.Println(err1)
	}

	// 3. application exit close listener port
	defer listener.Close()

	for {
		fmt.Println("start to connect")
		// 4. establish connection
		conn, err2 := listener.Accept()

		if err2 != nil {
			fmt.Println(err2)
		}

		// 5. bind RPC service
		rpc.ServeConn(conn)
	}

}
