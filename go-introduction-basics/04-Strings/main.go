package main

import (
	"fmt"
	"strings"
)

func main() {
	// slide 2: rune vs byte
	s := "élite"
	fmt.Printf("%8T %[1]v\n", s)
	fmt.Printf("%8T %[1]v\n", []rune(s))
	fmt.Printf("%8T %[1]v\n", []byte(s))

	// slide 3: same, in Chinese
	s = "你好 世界"
	fmt.Printf("%8T %[1]v\n", s)
	fmt.Printf("%8T %[1]v\n", []rune(s))
	fmt.Printf("%8T %[1]v\n", []byte(s))

	// slide 4: substrings share storage
	s = "hello, world"
	hello := s[:5]
	world := s[7:]
	fmt.Println(hello, world)

	// slide 5: len, slicing, +
	s = "the quick brown fox"
	a := len(s)                 // 19
	b := s[:3]                  // "the"
	c := s[4:9]                 // "quick"
	d := s[:4] + "slow" + s[9:] // replaces "quick"
	// s[5] = 'a'              // SYNTAX ERROR
	s += "es" // now plural (copied)
	fmt.Println(a, b, c, d, s)

	// slide 6: package strings
	s = "a string"
	x := len(s) // 8
	fmt.Println(x)
	fmt.Println(strings.Contains(s, "g"))   // true
	fmt.Println(strings.Contains(s, "x"))   // false
	fmt.Println(strings.HasPrefix(s, "a"))  // true
	fmt.Println(strings.Index(s, "string")) // 2
	s = strings.ToUpper(s)                  // "A STRING"
	fmt.Println(s)
}
