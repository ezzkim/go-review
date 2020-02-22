package main

import "fmt"

func main() {
	/*
		var a int
		var b bool
		var c string
		var d float32
		var e float64
	*/

	var (
		a int
		b bool
		c string
		d float32
		e float64
	)

	fmt.Printf("a=%d, b=%t, c=%s, d=%f, e=%f\n", a, b, c, d, e)

	a = 10
	b = true
	c = "hello"
	d = 4.8
	e = 1.234454
	fmt.Printf("a=%d, b=%t, c=%s, d=%f, e=%f\n", a, b, c, d, e)
}
