package main

import (
	"fmt"
	"time"
)

func main() {
	fortest()
	testmutilasign()
}

func fortest() {
	for i := 1; i <= 10; i++ {
		if i == 3 {
			continue // next loop
		}

		if i > 6 {
			break // end loop
		}

		fmt.Printf("i=%d\n", i)
	}

	fmt.Printf("-----------\n")

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Printf("\n")
	}

	fmt.Printf("-----------\n")

	i := 1
	//for ; i < 10; { // ok method, equals below
	for i < 10 { // ok too, more simple
		fmt.Printf("i=%d\n", i)
		i += 2
	}

	//for no,i:=10,1; i<10 && no<19; i+=1,no+=1 { // error i+=1,no+=1, must
	//for no:=10, i:=1; i < 10 && no < 19; i, no = i+1, no+1 { // error no:=10, i := 1;
	for no, i := 10, 1; i < 10 && no < 19; i, no = i+1, no+1 { //ok
		fmt.Printf("i=%d no=%d\n", i, no)
	}

	n := 1
	for {
		fmt.Printf("hello... %d\n", n)
		n++
		time.Sleep(1 * time.Second)
	}
}

func testmutilasign() {
	var a int
	var b int
	//a, b, = 10, 20
	a, b, c := 10, 20, "30" // asign mutil varies in one line
	fmt.Printf("a=%d b=%d c=%s\n", a, b, c)
}
