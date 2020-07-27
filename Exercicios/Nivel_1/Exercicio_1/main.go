package main

import (
	"fmt"
)

func main() {
	x, y, z := 42, "James Bond", true

	fmt.Printf("Valores: %d, %s, %t\n", x, y, z)
	fmt.Printf("1. %d\n", x)
	fmt.Printf("2. %s\n", y)
	fmt.Printf("3. %t\n", z)
}
