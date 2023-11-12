package main

import (
	"fmt"
	"math"
)

// go struct inherited concepts mostly follow c, can think of kind of equal to OOP object
// go doesn't support classes or inheritance in the complete sense
// there is also anonymous struct which dont have name and cannot be referenced other code
type messageToSend struct {
	number    int
	message   string
	sender    user
	recipient user
	// this is anonymous struct
	owner struct {
		name   string
		number int
	}
	// this car struct is embedded struct
	car
}

type user struct {
	name   string
	number int
}

type car struct {
	make  string
	model string
}

// go does support methods which can be defined on struct
type authenticationInfo struct {
	username string
	password string
}

func (authInfo authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: basic %s:%s", authInfo.username, authInfo.password)
}

// interface in go is collection of method signatures
type shape interface {
	area()
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	test(messageToSend{number: 123456, message: "thanks"})

}

func test(m messageToSend) {

	user := user{}
	user.name = "wong"
	user.number = 123456

	message1 := messageToSend{}
	message1.sender = user
	message1.recipient.name = "lee"
	message1.recipient.number = 654321

	user2 := messageToSend{
		number:  123456,
		message: "thanks",
		car: car{
			make:  "toyota",
			model: "camry",
		},
	}

	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.number)
	fmt.Println(user)
	fmt.Println(message1)
	fmt.Println(user2)

	// normally we think that should be access like user2.car.make
	// embedded struct fileds promotes to the top level so that we just directly access the embedded fields
	fmt.Println("user car: ", user2.make)
	fmt.Println("user car: ", user2.model)

	basicAuth := authenticationInfo{
		username: "wong",
		password: "123456",
	}
	fmt.Println(basicAuth.getBasicAuth())

}
