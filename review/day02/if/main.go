package main

import (
	"fmt"
)

func main() {
	testif()
}

func testif() {
	num := 10
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is even\n", num)
	}

	//if
	//else if
	//else if
	//else

	if num > 5 && num < 10 {
		fmt.Printf("5-10\n")
	} else if num >= 10 && num < 20 {
		fmt.Printf("10-20\n")
	} else if num >= 20 && num < 30 {
		fmt.Printf("20-30\n")
	} else {
		fmt.Printf("else\n")
	}

	//block inner
	if num2 := 20; num2%2 == 0 {
		fmt.Printf("%d is even\n", num2)
	}
	//fmt.Printf("%d", num2) //error scope
}
