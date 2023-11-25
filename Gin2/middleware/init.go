package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {

	fmt.Println(time.Now())

	fmt.Println(c.Request.URL)
	// middleware and controller share data
	c.Set("username", "james")

	// define goroutine for count log
	// for goroutine to work with middleware, cannot directly use *gin.Context, need to Copy its instance to use
	// normally goroutine will end once main coroutine exit and cloudn't executed finished here, but since go gin will alawys
	// up, so its coroutine never exit
	cCopy := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done in path " + cCopy.Request.URL.Path)
	}()
}
