package main

import (
	"fmt"
	"time"
)

// 定向通道的使用
func main() {

	//只写通道
	//ch1 := make(chan<- int)
	//ch1 <- 200
	//temp:=<-ch1

	//只读通道
	//ch2 := make(<-chan int)
	//ch2 <- 100

	ch := make(chan int)

	go writeOnly(ch)

	go readOnly(ch)

	time.Sleep(time.Second * 5)

}

func writeOnly(ch chan<- int) {
	ch <- 200
}

func readOnly(ch <-chan int) {
	data := <-ch
	fmt.Println(data)
}
