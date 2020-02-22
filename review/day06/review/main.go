package main

import (
	"fmt"
	"sort"
)

func main() {
	//show()
	//goSort()

	work3()
}

func show() {
	var sa = make([]string, 5, 10) // have 5 "" string
	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))
	}

	fmt.Printf("%v", sa)
}

func goSort() {
	var a [5]int = [5]int{5, 4, 3, 2, 1}
	sort.Ints(a[:])
	fmt.Printf("a=%v\n", a)

	ls := sort.IntSlice{
		1,
		4,
		5,
		3,
		2,
	}
	fmt.Println(ls) //[1 4 5 3 2]
	sort.Ints(ls)
	fmt.Println(ls) //[1 2 3 4 5]

	ls2 := sort.StringSlice{
		"100",
		"42",
		"41",
		"3",
		"2",
	}
	fmt.Println(ls2) //[100 42 41 3 2]
	sort.Strings(ls2)
	fmt.Println(ls2) //[100 2 3 41 42]

	ls3 := sort.StringSlice{
		"d",
		"ac",
		"c",
		"ab",
		"e",
	}
	fmt.Println(ls3) //[d ac c ab e]
	sort.Strings(ls3)
	fmt.Println(ls3) //[ab ac c d e]

	ls4 := sort.StringSlice{
		"啊",
		"博",
		"次",
		"得",
		"饿",
		"周",
	}
	fmt.Println(ls4) //[啊 博 次 得 饿 周]
	sort.Strings(ls4)
	fmt.Println(ls4) //[博 周 啊 得 次 饿]

	ls5 := sort.Float64Slice{
		1.1,
		4.4,
		5.5,
		3.3,
		2.2,
	}
	fmt.Println(ls5) //[1.1 4.4 5.5 3.3 2.2]
	sort.Float64s(ls5)
	fmt.Println(ls5) //[1.1 2.2 3.3 4.4 5.5]
}

func work3() {
	type testSlice [][]int

	func (l testSlice) Len() int      { return len(l) }
	func (l testSlice) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
	//func (l testSlice) Less(i, j int) bool { return l[i][1] < l[j][1] }
	func (l testSlice) Less(i, j int) bool { return l[i][0] < l[j][0] }

	ls := testSlice{
		{1, 4},
		{9, 3, 6},
		{7, 5},
	}

	fmt.Printf("len=%d", len(ls))

	fmt.Println(ls) //[[1 4] [9 3] [7 5]]
	sort.Sort(ls)
	fmt.Println(ls) //[[9 3] [1 4] [7 5]]
}

func work4() {
type testSlice []map[string]float64

func (l testSlice) Len() int           { return len(l) }
func (l testSlice) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l testSlice) Less(i, j int) bool { return l[i]["a"] < l[j]["a"] } //按照"a"对应的值排序


	ls := testSlice{
		{"a": 4, "b": 12},
		{"a": 3, "b": 11},
		{"a": 5, "b": 10},
	}

	fmt.Println(ls) //[map[a:4 b:12] map[a:3 b:11] map[a:5 b:10]]
	sort.Sort(ls)
	fmt.Println(ls) //[map[a:3 b:11] map[a:4 b:12] map[a:5 b:10]]
}

func work5() {
type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type testSlice []People

func (l testSlice) Len() int           { return len(l) }
func (l testSlice) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l testSlice) Less(i, j int) bool { return l[i].Age < l[j].Age }


	ls := testSlice{
		{Name: "n1", Age: 12},
		{Name: "n2", Age: 11},
		{Name: "n3", Age: 10},
	}

	fmt.Println(ls) //[{n1 12} {n2 11} {n3 10}]
	sort.Sort(ls)
	fmt.Println(ls) //[{n3 10} {n2 11} {n1 12}]
}

type People struct {
	Name string  `json:"name"`
	Age  float64 `json:"age"`
}

func isNaN(f float64) bool {
	return f != f
}

func work6() {
type testSlice []People

func (l testSlice) Len() int      { return len(l) }
func (l testSlice) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l testSlice) Less(i, j int) bool {
	return l[i].Age < l[j].Age || isNaN(l[i].Age) && !isNaN(l[j].Age)
}


	ls := testSlice{
		{Name: "n1", Age: 12.12},
		{Name: "n2", Age: 11.11},
		{Name: "n3", Age: 10.10},
	}

	fmt.Println(ls) //[{n1 12.12} {n2 11.11} {n3 10.1}]
	sort.Sort(ls)
	fmt.Println(ls) //[{n3 10.1} {n2 11.11} {n1 12.12}]
}
