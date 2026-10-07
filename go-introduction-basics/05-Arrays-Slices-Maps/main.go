package main

import "fmt"

func main() {
	arrays()
	slices()
	maps()
}

func arrays() {
	var a [3]int
	b := [3]int{0, 0, 0}
	c := [...]int{0, 0, 0} // sized by initializer

	var d [3]int
	d = b // elements copied

	m := [...]int{1, 2, 3, 4}
	// c = m // TYPE MISMATCH: [3]int vs [4]int

	fmt.Println("arrays:", a, b, c, d, m)
}

func slices() {
	var a []int      // nil, no storage
	b := []int{1, 2} // initialized

	a = append(a, 1) // append to nil OK
	b = append(b, 3) // []int{1, 2, 3}

	a = b // overwrites a

	d := make([]int, 5) // []int{0, 0, 0, 0, 0}
	e := a              // same storage (alias)

	t := []byte("string")
	fmt.Println("slices:", a, b, d, e, e[0] == b[0])
	fmt.Println(t[:2], t[2:], t[3:5]) // [115 116] [114 105 110 103] [105 110]
}

func maps() {
	var m map[string]int      // nil, no storage
	p := make(map[string]int) // non-nil but empty

	a := p["the"] // returns 0
	b := m["the"] // same thing
	// m["and"] = 1 // PANIC - nil map
	m = p
	m["and"]++    // OK, same map as p now
	c := p["and"] // returns 1

	v, ok := p["nope"] // 0, false
	fmt.Println("maps:", a, b, c, v, ok)
}
