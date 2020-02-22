package main

import (
	"fmt"
)

func sumsub(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	a := 11
	b := 16

	r1, r2 := sumsub(a, b)
	fmt.Printf("sum=%d, sub=%d\n", r1, r2)
}
