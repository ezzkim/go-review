package main

import "fmt"

func main() {
	forswitch()
	forswitch2()
	forswitch3()
	forswitch4()
	showtable()
}

//多值匹配用逗号分开
func forswitch() {
	//n := 3
	//switch n { // 0k
	//switch(n) { // ok

	//the switch block scope
	switch n := 2; n {
	case 1:
		fmt.Printf("is 1\n")
	//case 3 | 2: // bad, | is bit operator
	case 2, 3: // ok
		fmt.Printf("is 2 or 3\n")
	default:
		fmt.Printf("other value\n")
	}
}

//多值匹配用逗号分开
func forswitch2() {
	//letter := 'i'
	//switch letter {
	switch letter := 'i'; letter { //ok
	case 'a':
		fmt.Printf("is a\n")
	case 'b', 'c', 'i':
		//case 'b' | 'c' | 'i': // bad, | is bit operator
		fmt.Printf("is b , c, i\n")
	default:
		fmt.Printf("other letter\n")
	}
}

//switch use as if
func forswitch3() {
	n := 5
	switch {
	case n > 1 && n <= 3:
		fmt.Printf("in 1-3\n")
	case n > 4 && n <= 5:
		fmt.Printf("in 4-5\n")
	case n > 5 && n <= 10:
		fmt.Printf("in 5-10\n")
	default:
		fmt.Printf("no (1,10]\n")
	}
}

func number() int {
	num := 15 * 3
	return num
}

//switch 有两种用法
//1）switch 后面有变量， case 部分就用值来匹配这个变量， 都多用逗号
//2）switch 后无变量， case 后面是一个条件判断， 此时switch 用作if

//switch 后不可以是一个语句
func forswitch4() {
	//switch num := number(); num { // bad
	//switch num := number(); { //must add ;
	num := number()
	switch {
	case num < 50:
		fmt.Printf("%d less 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d less 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d less 200\n", num)
	}
}

func showtable() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
}
