package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Banana", "Orange"}

	fruits = append(fruits, "Mango") // adds the item at the end of the slice
	fruits[2] = "Cucumber"           // replaces the item at index 2

	for index, fruit := range fruits {
		fmt.Println(index, fruit)
	}
}
