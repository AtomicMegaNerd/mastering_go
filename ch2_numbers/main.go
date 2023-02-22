package main

import "fmt"

func main() {
	c1 := 12 + 1i       // We can create complex numbers this way
	c2 := complex(5, 7) // another way
	// In both cases on my machine it defaults to complex128, probably because this is a
	// 64-bit machine.

	// %T prints the type of something, interesting
	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of c2: %T\n", c2)

	// Some complex addition
	var c3 complex64 = complex64(c1 + c2)
	fmt.Println("c3: ", c3)
	fmt.Printf("Type of c3: %T\n", c3)

	cZero := c3 - c3
	fmt.Println("cZero:", cZero)

	x := 12
	k := 5

	fmt.Println(x)
	fmt.Printf("Type of x: %T\n", x)

	div := x / k
	fmt.Println("div:", div)

	// Remember Go initializes numbers to 0 if we do not specify a value
	var m, n float64
	m = 1.223
	fmt.Println("m, n:", m, n)

	y := 4 / 2.3
	fmt.Println("y:", y)

	divFloat := float64(x) / float64(k)
	fmt.Println("divFloat:", divFloat)
	fmt.Printf("Type of divFloat: %T\n", divFloat)

}
