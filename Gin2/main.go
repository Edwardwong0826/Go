package main

import (
	"Go/Gin2/routers"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Go Middleware
// router middleware is when access to router, it will go through and executed all the middleware
func initMiddleware1(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1. this is a middleware 1")

	// c.Next() this will invoke remaining handle program, means 1. -> initMiddleware -> /news func()
	// order of execution : 1. -> /news func() -> 2.
	// similiar to Java AOP around concepts
	c.Next()

	fmt.Println("2. this is a middleware 1")
	time.Sleep(time.Second)
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

func initMiddleware3(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1. this is a middleware 3")

	c.Next()

	fmt.Println("2. this is a middleware 3")
	//time.Sleep(time.Second)
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

func initMiddlewar2(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1. this is a middleware 2")

	// c.Abort() this will abort remaining handle program, means 1. -> initMiddleware2 -> skip /news func()
	// order of execution : 1. -> skip /news func() -> 2.
	// similiar to Java AOP around concepts
	c.Abort()

	fmt.Println("2. this is a middleware 2")
	time.Sleep(time.Second)
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

func initMiddleware4(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1. this is a middleware 4")

	c.Next()

	fmt.Println("2. this is a middleware 4")
	//time.Sleep(time.Second)
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

func main() {
	// create default router engine
	r := gin.Default()

	// Global Middleware
	r.Use(initMiddleware4)

	// all function not last, in the middle is called middleware
	// order of execution : initMiddleware1 ->  initMiddleware3 -> func()
	r.GET("/", initMiddleware1, initMiddleware3, func(c *gin.Context) {
		fmt.Println("this is a homepage")

		c.String(200, "Gin homepage")
	})

	r.GET("/news", initMiddlewar2, func(c *gin.Context) {

		c.String(200, "News homepage")
	})

	r.GET("/global", func(c *gin.Context) {

		c.String(200, "Global homepage")
	})

	// note - routers need to be small letter, if first letter is capital letter, will error
	// Routers name on folder can be capital letter
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)

	// start a web server
	// default start HTTP server at 0.0.0.0:8080 port
	r.Run(":8090")
}
