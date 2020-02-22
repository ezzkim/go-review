package main

import (
	"fmt"
	"time"
)

var length int = 50000
var arrays []int

func init() {
	for i := 0; i < length; i++ {
		arrays = append(arrays, length-i)
	}
}

func show(arrays []int, n int) {
	fmt.Printf("%v\n", arrays[:n])
}

func selectSort(arrays []int) {
	n := len(arrays)
	for i := 0; i < n-1; i++ {
		index := i
		for j := i + 1; j < n; j++ {
			if arrays[i] > arrays[j] {
				index = j
			}
		}
		if index != i {
			temp := arrays[i]
			arrays[i] = arrays[index]
			arrays[index] = temp
		}
	}
}

func bubbleSort(arrays []int) {
	var n int = len(arrays)
	var swap_flag bool

	for {
		swap_flag = false
		for i := 0; i < n-1; i++ {
			if arrays[i] > arrays[i+1] {
				temp := arrays[i]
				arrays[i] = arrays[i+1]
				arrays[i+1] = temp

				swap_flag = true
			}
		}
		if !swap_flag {
			break
		}
		n--
	}
}

func insertSort(arrays []int) {
	n := len(arrays)
	for i := 1; i < n; i++ {
		pos := i
		data := arrays[i]

		for j := i - 1; j >= 0; j-- {
			if arrays[j] > data {
				arrays[j+1] = arrays[j]
				pos--
			}
		}
		arrays[pos] = data
	}
}

func main() {
	show(arrays, 5)

	t1 := time.Now().Unix() // second
	//selectSort(arrays)
	//bubbleSort(arrays)
	insertSort(arrays)
	t2 := time.Now().Unix()
	show(arrays, 5)
	fmt.Printf("cost time:%d seconds\n", t2-t1)
}
