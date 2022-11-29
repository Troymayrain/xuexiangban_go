package main

import "fmt"

// channel的使用
func main() {
	//定义一个通道ch
	var ch chan bool
	//初始化ch
	ch = make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		//往通道写入数据
		ch <- true
	}()
	//从通道读取数据
	data := <-ch
	fmt.Println("读取通道的数据为：", data)

	//ch2 := make(chan int)
	//ch2 <- 12
}
