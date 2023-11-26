package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// if want to download third party dependency
// go mod init xxx first
// god mod tidy

type Goods struct{}

type AddGoodsRequest struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

type AddGoodsResponse struct {
	Success bool
	Message string
}

type GetGoodsRequest struct {
	Id int
}

type GetGoodsResponse struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

// Go had integrate its own RPC package, net and net/rpc is the one
// define one remote call function, function normally put into struct inside

// Go RPC rules
// 1. method only can have 2 able serialize parameter means cannot be channel and func type, and second parameter need to be pointer type
// 2. method need to return one error type, and must be public method, means first letter of method need to be capital letter
func (this Goods) AddGoods(request AddGoodsRequest, response *AddGoodsResponse) error {

	// 1. execute increment
	fmt.Println(request)
	// 2. return increment result
	*response = AddGoodsResponse{
		Success: true,
		Message: "execute success",
	}
	return nil
}

func (this Goods) GetGoods(request GetGoodsRequest, response *GetGoodsResponse) error {

	// 1. execute increment
	fmt.Println(request)
	// 2. return increment result
	*response = GetGoodsResponse{
		Id:      12,
		Title:   "title",
		Price:   24.5,
		Content: "content",
	}
	return nil
}

func main() {

	fmt.Println("Goods server")
	//1. Register RPC service
	err := rpc.RegisterName("goods", new(Goods))

	if err != nil {
		fmt.Println(err)
	}

	// 2. lister port
	listener, err1 := net.Listen("tcp", "127.0.0.1:8091")

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
