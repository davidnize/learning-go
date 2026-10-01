package main

import "fmt"

func main() {

	name := "David" // The `:=` syntax declares a new variable and lets Go infer its type from the value.
	age := 20       // A variable declared with `:=` can be assigned a new value later:
	age = 18

	fmt.Println("I am starting my Go journey!")
	fmt.Println("My name is", name)
	fmt.Println("I am", age, "years old.")
	fmt.Println("I am learning Go.")
}
