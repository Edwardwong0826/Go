package json

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title" xml:"title"`
	Desc    string `json:"desc" xml:"desc"`
	Content string `json:"content" xml:"content"`
}

type Article2 struct {
	// `` this symbol called backquote 反引号, for json body key to allow to have small letter
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func Json1(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "hello gin",
	})

}

func Json2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "hello gin",
	})

}

func Json3(c *gin.Context) {

	// dont know why, use this way to display JSON, type fields first letter required to be capital letter
	a := &Article2{
		Title:   "I am a title",
		Desc:    "description",
		Content: "test content",
	}

	c.JSON(http.StatusOK, a)

}

func Jsonp(c *gin.Context) {

	a := &Article{
		Title:   "I am a title",
		Desc:    "description",
		Content: "test content",
	}

	c.JSONP(http.StatusOK, a)

}
