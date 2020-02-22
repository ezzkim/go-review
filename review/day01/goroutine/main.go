package main

import (
	"fmt"
	"time"
)

func cacl() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("cacl finished\n")
}

func main() {
	go cacl()
	fmt.Printf("main finished\n")
	time.Sleep(11 * time.Second)
}
