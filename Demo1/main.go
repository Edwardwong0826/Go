package main

import "fmt"

var g = "global variable"

func main() {
	// go build main.go - command to build the executable file
	// .\main.exe - command to run the exe
	// go run main.go - build and run directly
	fmt.Println("Hello World")

	var a = "aaa"
	var b int = 10
	c := 30
	fmt.Printf("%v\n", a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%v %T\n", c, c)

	var (
		username string = "test"
		age      int    = 16
		sex      string = "M"
	)

	// := only can use declare local variable, not global variable
	fmt.Println(username, age, sex)

	// when there is multiple assignmnet and we to ignore some value, can use anonymous variable
	// anonymous variable use _, example
	var username2, _ = getUserinfo()
	fmt.Printf(username2)
	fmt.Println()

	const (
		c1 = iota // 0
		c2 = 100  // 100
		c3 = iota // 2
	)

	const c4 = iota
	fmt.Println(c1, c2, c3, c4)

}

func getUserinfo() (string, int) {
	return "test2", 10
}
