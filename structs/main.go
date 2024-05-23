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
	myPerson.changeName("Marc")

	a := 7
	b := &a // This is a pointer to `a`
	fmt.Println(b)
	*b = 9
	// This is a reference to the value of b, which is a's location. It's assigning 9 to the address 
	// that b points to, which is a's value, changing the value of a.
	// Asterisc is also doing what's called "Dereferencing" here, which is the process described above.
	// *b == a; b == &a
	fmt.Println(a)

	fmt.Printf("This is my new person %+v\n", myPerson)

	mySlice := []int {
		1, 2, 3,
	}

	// Using underscore instead of 'index' tells the compiler to ignore the value that's returned by range
	// since we're not going to use it. But, this only increments the copy, not the actual value
	// for _, value := range mySlice {
	// 	value++
	// }

	for index, _ := range mySlice {
		mySlice[index]++
	}

	fmt.Println(mySlice)
}