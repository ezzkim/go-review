package main

import (
	"fmt"
)

//const var_name type = value
//const var_name = value // must init

func main() {
	/*
		const a int = 123
		const b = "hello b"
		const c = true
	*/

	const (
		a int = 123
		b     = "hello"
		c     = true
	)

	const (
		aa int = 100
		bb     // = 100
		cc     // = 100
		dd = 200
		ee // = 200
	)

	fmt.Printf("a=%d, b=%s, c=%t\n", a, b, c)
	fmt.Printf("aa=%d, bb=%d, cc=%d, dd=%d, ee=%d\n", aa, bb, cc, dd, ee)

	//iota auto increase each assign and init to 0 when face const
	const (
		e = iota
		f // = iota
		g // = iota
	)
	fmt.Printf("e=%d, f=%d, g=%d\n", e, f, g)

	const ( //iota = 0
		a0 = 1         // iota += 1
		a1 = 1 << iota // iota += 1
		a2             // = 1 << iota
		a3             // = 1 << iota
	)

	fmt.Printf("a0=%d, a1=%d, a2=%d, a3=%d\n", a0, a1, a2, a3)

	testIota()
}

func testIota() {
	const (
		A = iota
		B
		C
		D = 8
		E
		F = iota
		G
	)

	fmt.Printf("%d,%d,%d,%d,%d,%d,%d\n",A,B,C,D,E,F,G)
}