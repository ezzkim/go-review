package main

import (
	"fmt"

	"ericsson.com/study/review/day04/calc"
)

func main() {
	s1 := 200
	s2 := 300
	sum := calc.Add(s1, s2)
	fmt.Printf("sum=%d\n", sum)

	//sb := calc.sub(a,b) // can not access outside package

	fmt.Printf("calc.A=%d\n", calc.A)
	//fmt.Printf("calc.a=%d\n", calc.a) // can not access outside package

	fmt.Printf("sub=%d\n", calc.Test()) //Test can access inside package
}
