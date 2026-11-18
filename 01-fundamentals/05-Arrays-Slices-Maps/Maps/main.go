package main

import "fmt"

func main() {
	inventory := map[string]int{
		"Mouse": 1,
		"Cable": 3,
	}
	inventory["laptop"] = 1
	inventory["keyboard"] = 1

	inventory["laptop"] += 1

	quantity, ok := inventory["laptop"]

	if ok {
		fmt.Println("Quantity:", quantity)
	} else {
		fmt.Println("Item not found")
	}

	for item, quantity := range inventory {
		fmt.Println(item, quantity)
	}

	fmt.Println(inventory)
}
