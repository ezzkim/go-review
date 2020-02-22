package main

import (
	"fmt"
	"math/rand"
	"time"
)

func work1(a []int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	/*
		for _, item := range a {
			sum += item
		}
	*/
	return sum
}

func work2(a []int, sum int) [][2]int {
	var rest [][2]int

	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if a[i]+a[j] == sum {
				one := [2]int{i, j}
				rest = append(rest, one)
			}
		}
	}

	return rest
}

func main() {
	//var a [3]int = [3]int{1,2,3}
	//var a [3]int = [3]int{1:33}
	//a := [...]int{3:12}

	a := [...]int{1, 2, 3, 4, 5}

	s := work1(a[:])
	fmt.Printf("sum=%d\n", s)

	fmt.Printf("--------------\n")

	b := [...]int{1, 3, 5, 8, 7}
	rest := work2(b[:], 8)
	for _, item := range rest {
		fmt.Printf("index (%d, %d) sum = 8\n", item[0], item[1])
	}

	fmt.Printf("--------------\n")

	randTest()
}

func randTest() {
	rand.Seed(time.Now().Unix())

	var b [10]int
	for i := 0; i < 10; i++ {
		b[i] = rand.Intn(1000)
	}

	fmt.Printf("%v\n", b)

	sum := 0
	for _, item := range b {
		sum += item
	}
	fmt.Printf("sum=%d\n", sum)
}
