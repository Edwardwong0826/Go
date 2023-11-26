package main

import (
	"fmt"
	"net/rpc"
)

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

func main() {

	fmt.Println("client")

	//. 1 use rpc.dial connect with rpc service
	conn, err := rpc.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println(err)
	}

	//2. when client exit close the connection
	defer conn.Close()

	// 3. innvoke remote function
	var reply string
	// conn.Call first parameter refer to register name and the function
	err2 := conn.Call("hello.SayHello", "this is client", &reply)
	if err2 != nil {
		fmt.Println(err2)
	}

	// 4. get RPC return data
	fmt.Println(reply)

	conn2, err3 := rpc.Dial("tcp", "127.0.0.1:8091")
	if err3 != nil {
		fmt.Println(err3)
	}

	defer conn2.Close()

	var reply2 AddGoodsResponse
	err4 := conn2.Call("goods.AddGoods", AddGoodsRequest{
		Id:      10,
		Title:   "title",
		Price:   23.5,
		Content: "content",
	}, &reply2)

	if err4 != nil {
		fmt.Println(err4)
	}

	var reply3 GetGoodsResponse
	err5 := conn2.Call("goods.GetGoods", GetGoodsRequest{
		Id: 12,
	}, &reply3)

	if err4 != nil {
		fmt.Println(err5)
	}

	fmt.Println(reply3)

}
