package main

import "fmt"

func main() {
	students := map[string]map[string]int{
		"David": {
			"Math":  90,
			"Addma": 100,
		},
		"Nize": {
			"Math":  100,
			"Addma": 100,
		},
	}
	for name, subject := range students {
		fmt.Println("Student:", name)

		for subject, mark := range subject {
			fmt.Println(subject, mark)
		}
	}
}
