package main

import (
	"fmt"
)

func testArray() {
	var a [3]int // length is the part of array
	fmt.Printf("a=%+v\n", a)

	//define and the init
	a[0] = 200
	a[2] = 300
	fmt.Printf("a=%+v\n", a)

	//a[3] = 500 // overflow
	//fmt.Printf("a=%+v\n", a)

	////////////////
	//define and init full data
	var a1 [3]int = [3]int{1, 2, 3}
	fmt.Printf("a1=%+v\n", a1)

	//define and init partion element
	a2 := [3]int{3, 2}
	fmt.Printf("a2=%+v\n", a2)

	//define and init full, skip show the length
	a3 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("a3=%+v\n", a3)

	//define
	a4 := [...]int{2: 8}       // only init index 2 value 8
	fmt.Printf("a4=%+v\n", a4) //a4=[0 0 8]

	a5 := [5]int{1: 100, 3: 300}
	fmt.Printf("a5=%+v\n", a5)

	a6 := [...]int{1: 100, 3: 300} //a6=[0 100 0 300]
	fmt.Printf("a6=%+v\n", a6)

	//length is the self partion of array
	fmt.Printf("a4:%T, a5：%T\n", a4, a5)
}

func testArray2() {
	a6 := [...]int{1: 100, 3: 300} //a6=[0 100 0 300]
	fmt.Printf("a6=%+v\n", a6)
	fmt.Printf("length=%d\n", len(a6))

	n := len(a6)
	for i := 0; i < n; i++ {
		fmt.Printf("a6[%d]=%d\n", i, a6[i])
	}

	for i, elem := range a6 {
		fmt.Printf("%d : %d\n", i, elem)
	}
}

func testArray11() {
	var a [3][2]int
	a[0][0] = 10
	a[0][1] = 20
	a[1][0] = 30
	a[1][1] = 30
	a[2][0] = 30
	a[2][1] = 30

	//fmt.Println(a)
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
	fmt.Println("other method")
	for i, val := range a {
		fmt.Printf("row[%d]=%v\n", i, val)
		for j, val2 := range val {
			fmt.Printf("(%d,%d)=%d ", i, j, val2)
		}
		fmt.Println()
	}
}

//array is value type, copy is value copy
func testArray12() {
	a := [3]int{10, 20, 30}
	b := a
	b[0] = 1000
	fmt.Printf("a=%v\n", a)
	fmt.Printf("b=%v\n", b)
}

//value copy
func testArray13() {
	var a int = 1000
	b := a
	b = 3000
	fmt.Printf("a=%d b=%d\n", a, b)
}

//parameter is slice
func show(a []int) {
	for i, v := range a {
		fmt.Printf("a[%d] : %\nd", i, v)
	}
	fmt.Printf("-----end----\n")
}

func showtest() {
	var a1 [3]int = [3]int{1, 2}
	show(a1[:]) // convert array to slice

	a2 := [5]int{1, 2, 3}
	show(a2[:])

	a3 := [...]int{1, 2, 3, 4, 6}
	show(a3[:])

	a4 := [...]int{2: 4, 5: 12}
	show(a4[:])
}

////////////

//slice is a reference
func modifySlice(b []int) {
	b[0] = 111
}

func testArray14() {
	var a [3]int = [3]int{10, 20, 30}
	//modify(a)
	modifySlice(a[:]) // a[:] is a slice
	fmt.Println(a)
}

func main() {
	//testArray()
	//testArray2()
	//testArray11()
	//testArray12()
	testArray13()
	//showtest()
	testArray14()
}
