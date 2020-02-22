package main

import (
	"fmt"
	"unicode"
)

func main() {
	//arrs := work1()
	//arrs := work2()
	//fmt.Printf("%v\n", arrs)
	var str string = "hello 你好中国!#%$1234 56"
	cn, nn, sn, on := work3(str)
	fmt.Printf("str:%s\n", str)
	fmt.Printf("char:%d, num:%d, spac:%d, other:%d\n", cn, nn, sn, on)
	cn, nn, sn, on = work4(str)
	fmt.Printf("char:%d, num:%d, spac:%d, other:%d\n", cn, nn, sn, on)
}

func justfy(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

//all the oven data between 1:100
func work1() []int {
	var arrs []int
	for i := 2; i < 100; i++ {
		if justfy(i) {
			arrs = append(arrs, i)
		}
	}

	return arrs
}

func isShuixian(n int) bool {
	g := n % 10
	s := (n / 10) % 10
	b := n / 100

	number := g*g*g + s*s*s + b*b*b
	return n == number
}

//get all shuixian data between 100-1000
func work2() []int {
	var arrs []int
	for i := 100; i < 1000; i++ {
		if isShuixian(i) {
			arrs = append(arrs, i)
		}
	}

	return arrs
}

/*
包 unicode 包中一些常用函数（其中 v 代表字符）：

判断是否为字母： unicode.IsLetter(v)
判断是否为十进制数字： unicode.IsDigit(v)
判断是否为数字： unicode.IsNumber(v)
判断是否为空白符号： unicode.IsSpace(v)
判断是否为Unicode标点字符 :unicode.IsPunct(v)
*/

func work4(str string) (charCount int, numCount int, spaceCount int, otherCount int) {
	utfChars := []rune(str)
	//unicode.IsGraphic(r)

	for _, v := range utfChars {
		switch {
		case unicode.IsLetter(v): //!#%$ is letter
		//case unicode.IsGraphic(v):
			charCount++
		case unicode.IsNumber(v):
			numCount++
		case unicode.IsSpace(v):
			spaceCount++
		default:
			otherCount++
		}
	}

	return
}

func work3(str string) (charCount int, numCount int, spaceCount int, otherCount int) {
	utfChars := []rune(str) //consider chinese char
	n := len(utfChars)

	for i := 0; i < n; i++ {
		if (utfChars[i] >= 'a' && utfChars[i] <= 'z') || (utfChars[i] >= 'A' && utfChars[i] <= 'Z') {
			charCount++
			continue
		}

		if utfChars[i] >= '0' && utfChars[i] <= '9' {
			numCount++
			continue
		}

		if utfChars[i] == ' ' {
			spaceCount++
			continue
		}

		otherCount++
	}

	return
}
