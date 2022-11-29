package main

import (
	"fmt"
	"time"
)

// select使用
func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- 100
	}()
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- 200
	}()

	select {
	case num1 := <-ch1:
		fmt.Println("ch1:", num1)
	case num2 := <-ch2:
		fmt.Println("ch2:", num2)
		//default:
		//	fmt.Println("没有匹配到！")

	}
}
