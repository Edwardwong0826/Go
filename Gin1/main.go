package main

import (
	routers "Go/Gin1/routers"

	"github.com/gin-gonic/gin"
)

// type Article struct {
// 	Title   string `json:"title" xml:"title"`
// 	Desc    string `json:"desc" xml:"desc"`
// 	Content string `json:"content" xml:"content"`
// }

// type Article2 struct {
// 	// `` this symbol called backquote 反引号, for json body key to allow to have small letter
// 	Title   string `json:"title"`
// 	Desc    string `json:"desc"`
// 	Content string `json:"content"`
// }

// type User struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type UserInfo struct {
// 	Username string `json:"username" form:"username"`
// 	Password string `json:"password" form:"password"`
// }

// use below command to build and run at one time, else just run it will prompt firewall everytime
// go build main.go && .\main.exe

// for hot reload, there is two library, use the first one
// once install air, just run air
// air - https://github.com/cosmtrek/air
// gravityblast - https://github.com/gravityblast/fresh

// currently air hot reload not support well debugger
// when debug in VS code, just go click left section RUN AND DEBUG, click this Launch Package button
// https://superuser.com/questions/1785007/can-i-safely-disable-a-firewall-rule-added-by-visual-studio-code
// if keep block by windows firewall, go to Windows Defender Firewall with Advanced Security
// go to Inbound Rules -> New Rule -> Port -> put 8080 port -> and next

// https://maelvls.dev/go111module-everywhere/
// when get if there is error about GO111MODULE=off
// in windows - set GO111MODULE=on
// so that can use go install xxx
func main() {
	// create default router engine
	r := gin.Default()

	// note - routers need to be small letter, if first letter is capital letter, will error
	// Routers name on folder can be capital letter
	routers.ApiRoutersInit(r)
	routers.OriginRoutersInit(r)
	routers.JsonRoutersInit(r)
	routers.XmlRoutersInit(r)

	// start a web server
	// default start HTTP server at 0.0.0.0:8080 port
	r.Run(":8080")
}
