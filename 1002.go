package main

import (
	"fmt"
)

func main() {
	var a float64
	fmt.Scan(&a)
	n := 3.14159
	area := n * (a * a)
	fmt.Printf("A=%.4f", area)
	fmt.Println()
}
