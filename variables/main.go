package main

import "fmt"

func main() {
	// var myName string = "Drew"
	myName := "Drew"
	myInt := 10
	myFloat := 10.0

	// fmt.Printf("Hello, my name is %s\n", myName)
	fmt.Printf("Hello, my name is still %s, my int is %d, and my float is %f\n", myName, myInt, myFloat)

	var myFriendsName string
	var myBool bool
	var myOtherInt int

	myFriendsName = "Nick"

	fmt.Printf("my other friends name %s my bool %t and my other int %d\n", myFriendsName, myBool, myOtherInt)
}