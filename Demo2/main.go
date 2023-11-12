package main

import (
	"errors"
	"fmt"
)

var g = "global variable"

func main() {
	// go build main.go - command to build the executable file
	// .\main.exe - command to run the exe
	// go run main.go - build and run directly
	// dont know why, need ctls + s and run only will display latest changes on console
	fmt.Println("Hello World")

	// in Go, variable are pass by value not pass by reference
	basicVariables()
	constants()
	formatting()
	fmt.Println(concat("Go", " is used to developed EVM"))

	x1 := 10
	x2 := 20
	passByValue(x1, x2)
	fmt.Println(x1, x2)

	// we can igore the return value that we don't want, put _ to explicitly ignore variables
	// Go does not allow to have unused declare variables
	x3, _ := functionMultipleRetunValue()
	fmt.Println(x3)

	fmt.Println(earlyReturns(0, 2))
	fmt.Println(earlyReturns(10, 2))
}

func getUserinfo() (string, int) {
	return "test2", 10
}

func basicVariables() {
	// bool
	// string
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte - alias for uint8
	// rune - alias for uint32, represents a unicode code point
	// float32 float62
	// complex64 compltx128

	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	congrats := "happy birthday"
	pennisPerText := 2
	mileage, company := 80276, "Tesla"

	// in go we can do downcasting or upcasting, be mind of losing percision when downcasting
	accountAge := 2.6
	accountAgeInt := int(accountAge)

	fmt.Println("-------------")
	fmt.Println(congrats)
	//%T(need to be capital) this can put in Printf to know the variable type
	//%\n add extra space in Printf
	fmt.Printf("this variable is type %T\n", pennisPerText)
	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
	fmt.Println(mileage, company)
	fmt.Println(accountAgeInt)
}

func constants() {
	const const1 = "plan 1"
	const const2 = "plan 2"

	fmt.Println("const 1", const1)
}

func formatting() {
	// in fmt.Printf()
	// %s - interpolate as string
	// %d - interpolate an integer in decimal form
	// %f - interpolate a decmial
	fmt.Printf("I am %.2f\n", 10.523)
}

// this is example of function signature
// if all of the function parameter are same type, the type only needs to be declared at the last one
// func concat(s1, s2 string) string{...}
func concat(s1 string, s2 string) string {

	return s1 + s2
}

func functionMultipleRetunValue() (string, string) {
	return "Go", "ETH"
}

func getCoords() (x, y int) {
	// x and y are initialized with zero values
	// Go can return without named return values, this is know as naked return or implicit return
	// naked return statements should be used in short function, use in longer functions harm readbility
	return // automatically returns x and y

	// above is same as
	// func getCoords() (int, int){
	// var x int
	// var y int
	// return x, y
	// }

	// func getCoords() (x, y int){
	// 	return x, y - this is explicit
	// }

	// func getCoords() (x, y int){
	// 	return 5, 6 - this is explicit
	// }

	// func getCoords() (x, y int){
	// 	return - this is explicit
	// }

}

// Go support return early from a function, good to use in guard clauses
// guard clause is an early return from a function when a given condition is met
// instead of make nested conditionals one-dimensional (if/else chains )
func earlyReturns(x1 int, x2 int) (int, error) {
	if x1 == 0 {
		return 0, errors.New("Can't divide by zero")
	}

	return x1 / x2, nil
}

func passByValue(x1 int, x2 int) {
	x1 = x1 + x2
	fmt.Println(x1)
}
