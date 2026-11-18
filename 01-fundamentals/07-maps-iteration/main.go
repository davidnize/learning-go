package main

import "fmt"

func main() {
	students := map[string]map[string]int{
		"David": {
			"Math":    100,
			"Physics": 90,
		},
		"Nize": {
			"Math":    100,
			"Physics": 85,
		},
		"Shalom": {
			"Math":  87,
			"Addma": 100,
		},
	}
	for name, subjects := range students {
		fmt.Println("student:", name)

		for subject, mark := range subjects {
			fmt.Println(subject, mark)
		}
	}
}

// Alternative METHOD

// func main() {
// 	students := make(map[string]map[string]int)

// 	students["David"] = make(map[string]int)
// 	students["Nize"] = make(map[string]int)
// 	students["Shalom"] = make(map[string]int)

// 	students["David"]["Math"] = 100
// 	students["David"]["Physics"] = 90

// 	students["Nize"]["Math"] = 100
// 	students["Nize"]["Physics"] = 85

// 	students["Shalom"]["Math"] = 87
// 	students["Shalom"]["Physics"] = 100

// 	for name, subjects := range students {
// 		fmt.Println("Student:", name)

// 		for subject, mark := range subjects {
// 			fmt.Println(subject, mark)
// 		}
// 	}
// 	// fmt.Println(students)
// }
