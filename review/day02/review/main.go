package main

import (
	"fmt"
	"time"
)

func main() {
	timeTest()
	testcost()
}

func timeTest() {
	now := time.Now()
	//"2006/01/02 15:04:05" go birth time
	timestr := now.Format("2006/01/02 15:04:05")
	fmt.Printf("time:%s\n", timestr)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	timestr2 := fmt.Sprintf("%04d/%02d/%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
	fmt.Printf("time:%s\n", timestr2)
}

func testcost() {
	//t1 := time.Now().Unix() // the second from 1970.01.01 00:00:00
	t1 := time.Now().UnixNano() // the na second from
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond) // 1 millisecond
	}
	t2 := time.Now().UnixNano()
	cost := (t2 - t1) / 1000
	fmt.Printf("cost=%d ms\n", cost)
}
