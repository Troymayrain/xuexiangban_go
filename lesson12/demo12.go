package main

import (
	"fmt"
	"time"
)

// 定时器timer的使用
func main() {

	timer := time.NewTimer(time.Second * 3)
	fmt.Println(time.Now())
	ch := timer.C
	fmt.Println(<-ch)

	after := time.After(time.Second * 5)
	fmt.Println(time.Now())
	fmt.Println(<-after)

	time.AfterFunc(time.Second*2, test5)

	time.Sleep(time.Second * 10)
}

func test5() {
	fmt.Println("aaaaa")
}
