package main

import "fmt"

func main() {

	type student struct {
		Age       int
		Math      int
		Physics   int
		Chemistry int
	}

	students := map[string]student{
		"David": {
			Age:       18,
			Math:      100,
			Physics:   100,
			Chemistry: 98,
		},
	}
	students["Shalom"] = student{
		Age:       18,
		Math:      87,
		Physics:   100,
		Chemistry: 95,
	}

	david := students["David"]
	david.Math = 97
	students["David"] = david

	fmt.Println(students)

	for name, details := range students {
		fmt.Println("Student:", name)
		fmt.Println("Age:", details.Age)
		fmt.Println("Math:", details.Math)
		fmt.Println("Physics:", details.Physics)
		fmt.Println("Chemistry:", details.Chemistry)
	}
}
