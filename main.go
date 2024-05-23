package main

import (
	"fmt"
	"goFmBasics/circles"
	"goFmBasics/imports"
)

func main() {
	fmt.Println("hello world")

	newTicket := imports.Ticket {
		ID: 123,
		Event: "FEM course",
	}

	newTicket.PrintEvent()
	fmt.Println(newTicket)

	// Exercise for first module
	myCircle := circles.NewCircle(4.567)
	fmt.Println("The circle's circumference is ", myCircle.GetCircumference())
	fmt.Println("The circle's area is ", myCircle.GetArea())
}