package main

import (
	"fmt"
)

func main() {
	//testdefer()
	testdefer2()
}

func testdefer() {
	defer fmt.Printf("defer1\n")
	defer fmt.Printf("defer2\n")
	defer fmt.Printf("defer3\n")

	fmt.Printf("aaa\n")
	fmt.Printf("bbb\n")
}

func testdefer2() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("defer i=%d\n", i)
	}

	fmt.Printf("running\n")
	fmt.Printf("return\n")
}

//why defer i=0
func testDefer3() {
	var i int = 0
	defer fmt.Printf("defer i=%d\n", i)
	i = 1000
	fmt.Printf("i=%d\n", i)
}
