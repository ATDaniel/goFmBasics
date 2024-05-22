package main

import "fmt"

type Person struct {
	Name string
	Age int
}

// Unnecessary, but good practice to use a helper function
func NewPerson(name string, age int) Person {
	return Person {
		Name: name,
		Age: age,
	}
}

// This is called a Reciever. Takes in a reference (asterisc) before declaration.
// Acts similar to a class method in OOP
func (p *Person) changeName(newName string) {
	fmt.Println("address of copy", &p.Name)
	p.Name = newName
}

func main() {
	// if you don't initialize all values, it defaults to the zero value
	// myPerson := Person {
	// 	Name: "Drew",
	// }

	myPerson := NewPerson("Drew", 31)
	fmt.Println("address of allocated memory", &myPerson.Name)
	// fmt.Printf("This is my person %+v\n", myPerson)

	myPerson.changeName("Marc")

	fmt.Printf("This is my new person %+v", myPerson)
}