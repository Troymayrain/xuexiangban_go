package main

import (
	"fmt"
	"strconv"
	"time"
)

// 缓冲通道的使用
func main() {

	ch1 := make(chan int)
	fmt.Println(len(ch1), cap(ch1))
	//定义缓冲通道
	ch2 := make(chan int, 5)
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 2
	fmt.Println(len(ch2), cap(ch2))

	//定义缓冲通道
	ch3 := make(chan string, 10)

	//写入通道
	go test4(ch3)

	//读取通道
	for v := range ch3 {
		time.Sleep(time.Second)
		fmt.Println(v)
	}

}

func test4(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "test" + strconv.Itoa(i)
		fmt.Println("通道写入的数据：", "test"+strconv.Itoa(i))
	}
	close(ch)
}
