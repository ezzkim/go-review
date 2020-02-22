package main

import (
	"fmt"
)

func testMake1() {
	var a []int
	//如下操作create一个切片，初始5个元素，最大容量10， 5个元素的初始化是0
	a = make([]int, 5, 10) // init 5 element, can the capacity is 10
	a[0] = 10
	//a[1] = 20
	//a[5] = 50 // will panic, because the slice len is 5, max index is 4
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	a = append(a, 11)
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))

	for i := 0; i < 8; i++ {
		a = append(a, i)
		fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	}
	//观察切片的扩容操作，扩容的策略是翻倍扩容
	a = append(a, 1000)
	fmt.Printf("扩容之后的地址：a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
}

func testMake2() {
	var a []int
	//a = make([]int, 5, 10)
	a = make([]int, 5)
	//a[5] = 100
	fmt.Printf("a=%v len:%d cap:%d\n", a, len(a), cap(a))
	a = append(a, 10)
	//fmt.Printf("a=%v\n", a)
	fmt.Printf("a=%v len:%d cap:%d\n", a, len(a), cap(a))

	b := make([]int, 0, 10)
	fmt.Printf("b=%v len:%d cap:%d\n", b, len(b), cap(b))
	b = append(b, 100)
	fmt.Printf("b=%v len:%d cap:%d\n", b, len(b), cap(b))

	c := make([]int, 0)
	fmt.Printf("c=%v len:%d cap:%d\n", c, len(c), cap(c))
	c = append(c, 100)
	fmt.Printf("c=%v len:%d cap:%d\n", c, len(c), cap(c))
	c = append(c, 200)
	fmt.Printf("c=%v len:%d cap:%d\n", c, len(c), cap(c))
}

func testCap() {
	a := [...]string{"a", "b", "c", "d", "d", "f", "g", "h", "i"}
	b := a[1:3]
	fmt.Printf("b:%v len:%d cap:%d\n", b, len(b), cap(b)) // will cap is 8, because b is a's reference
}

func testCap2() {
	a := [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	b := a[1:3]
	fmt.Printf("b:%v len:%d cap:%d\n", b, len(b), cap(b))
	b = b[:cap(b)] // reslice
	fmt.Printf("after reslice :b:%v len:%d cap:%d\n", b, len(b), cap(b))

	fmt.Printf("a-addr=%p, b-addr=%p\n", &a, b)  //b form a[1]
	fmt.Printf("a[1]=%p, b-addr=%p\n", &a[1], b) //b form a[1]
	b = append(b, "hello")                       // re-melloc the memory
	fmt.Printf("a-addr=%p, b-addr=%p\n", &a, b)
}

func testSlice() {
	var a []int
	fmt.Printf("%p， len:%d, cap:%d\n", a, len(a), cap(a))
	if a == nil {
		fmt.Printf("a is nil\n")
	}

	a = append(a, 100)
	fmt.Printf("%p， len:%d, cap:%d\n", a, len(a), cap(a))
	a = append(a, 200)
	fmt.Printf("%p， len:%d, cap:%d\n", a, len(a), cap(a))
	a = append(a, 300)
	fmt.Printf("%p， len:%d, cap:%d\n", a, len(a), cap(a))
	a = append(a, 400)
	fmt.Printf("%p， len:%d, cap:%d\n", a, len(a), cap(a))
	a = append(a, 500)
	fmt.Printf("%p， len:%d, cap:%d\n", a, len(a), cap(a))
}

func testAppend() {
	var a []int = []int{1, 3, 4}
	var b []int = []int{4, 5, 6}

	a = append(a, 23, 34, 45)
	fmt.Printf("a=%v\n", a)
	a = append(a, b...) // b ... extend the b slice contents
	fmt.Printf("a=%v\n", a)
}

//a[:] : array convert to slice
//a... : extends slice or array

// the parameter is a slice, not array, array must contain the length
func sumArray(a []int) int {
	var sum int = 0
	for _, v := range a {
		sum = sum + v
	}
	return sum
}

func testSumArray() {
	var a [10]int = [10]int{1, 3, 3, 4, 5, 5, 8}
	sum := sumArray(a[:])
	fmt.Println("sum:", sum)

	var b [3]int = [3]int{10, 20, 30}
	sum = sumArray(b[:])
	fmt.Println("sum:", sum)
}

func modifySlice(a []int) {
	a[0] = 1000
}

func testModifySlice() {
	var a [3]int = [3]int{1, 2, 3}
	modifySlice(a[:])
	fmt.Println(a)
}

func testModifyCopy() {
	var a []int = []int{1}
	var b []int = []int{4, 5, 6}
	fmt.Printf("a=%v, a-addr=%p, len(a)=%d, cap(a)=%d\n", a, a, len(a), cap(a))
	fmt.Printf("b=%v, b-addr=%p, len(b)=%d, cap(b)=%d\n", b, b, len(b), cap(b))
	copy(a, b) // copy is go base function, copy 只是内容复制不会扩容切片
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Printf("a=%v, a-addr=%p, len(a)=%d, cap(a)=%d\n", a, a, len(a), cap(a))
	fmt.Printf("b=%v, b-addr=%p, len(b)=%d, cap(b)=%d\n", b, b, len(b), cap(b))
}

func main() {
	//testMake1()
	//testMake2()
	//testCap()
	//testCap2()
	//testSlice()
	//testAppend()
	//testSumArray()
	//testModifySlice()
	testModifyCopy()
}
