package main

import (
	"fmt"
)

func sayHello() {
	fmt.Printf("hello\n")
}

//func add(a int, b int) int {
func add(a, b int) int {
	return a + b
}

//return mutil value
func calc(a, b int) (int, int) {
	return a + b, a - b
}

//sum, sub have been defined
func calc2(a, b int) (sum int, sub int) {
	sum, sub = a+b, a-b
	return
}

//0 to more
func calc_v1(b ...int) int {
	sum := 0
	for _, item := range b {
		sum += item
	}

	return sum
}

//1 to more
func calc_v2(a int, b ...int) int {
	sum := a
	for _, item := range b {
		sum += item
	}

	return sum
}

//2 to more
func calc_v3(a, b int, c ...int) int {
	sum := a + b
	for _, item := range c {
		sum += item
	}

	return sum
}

func main() {
	sayHello()

	sum, _ := calc2(23, 5)
	fmt.Printf("sum:%d\n", sum)

	fmt.Printf("%d\n", calc_v1())
	fmt.Printf("%d\n", calc_v1(1, 2))
	fmt.Printf("%d\n", calc_v1(1, 2, 3))

	fmt.Printf("----------\n")

	fmt.Printf("%d\n", calc_v2(2))
	fmt.Printf("%d\n", calc_v2(2, 4))
}
