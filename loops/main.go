package main

import (
	"fmt"
	"slices"
)

func main() {
	// arrays
	// animals := [2]string{}

	animalss := [2]string{
		"dog",
		"cat",
	}

	animalss[0] = "dog"
	animalss[1] = "cat"

	fmt.Println(animalss)


	// Slice - dynamic lenght array
	animals := []string{
		"dog",
		"cat",
	}

	animals = append(animals, "moose")
	fmt.Println(animals)

	// removing an element from a slice
	animals = slices.Delete(animals, 0, 1) // slice, startRange, endRange
	fmt.Println(animals)
	animals = append(animals, "moose")

	// Old way of doing it
	// for i:=0; i < len(animals); i++ {
	// 	fmt.Printf("This is my animal: %s\n", animals[i])
	// }

	// how to do it in go 1.22+
	for index, value := range animals {
		fmt.Printf("This is my index %d and this is my animal %s\n", index, value)
	}

	for value := range 10 {
		fmt.Println(value);
	}

	// No while loops in Go. This is essentially the same
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++;
	}
}