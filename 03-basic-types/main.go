package main

import "fmt"

func main() {
	var name string = "David" // A `string` contains text. Strings use double quotes for ordinary string literals.
	var age int = 18
	var height float64 = 1.75  // `float64` is commonly used when you need a floating-point value.
	var isLearning bool = true // A `bool` represents either `true` or `false`.
	var weight float32         // When a variable is declared without an explicit value, Go gives it a zero value.

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Weight:", weight)
	fmt.Println("Learning Go:", isLearning)

	// Explicit type conversion.
	ageAsFloat := float64(age)
	fmt.Println("Age as float64:", ageAsFloat)

	// Print the types.
	fmt.Printf("name: %T\n", name) // The `%T` verb in `fmt.Printf` prints the type of the variable.
	fmt.Printf("age: %T\n", age)
	fmt.Printf("height: %T\n", height)
	fmt.Printf("isLearning: %T\n", isLearning)
}
