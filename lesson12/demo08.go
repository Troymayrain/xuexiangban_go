package main

import (
	"fmt"
	"time"
)

// 不要以共享内存的方式来通信，而要以通信的方式来共享内存
func main() {
	ch := make(chan int)
	//往通道写入数据
	go test3(ch)

	//fori读取通道数据
	/*for {
		time.Sleep(time.Second)
		v, ok := <-ch
		if !ok {
			fmt.Println("数据读取完毕，通道已关闭")
			break
		}
		fmt.Println("读取到的数据为：", v)
	}*/

	//for-range读取通道数据
	for v := range ch {
		time.Sleep(time.Second)
		fmt.Println("读取到的数据为：", v)
	}

}

func test3(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	//写完数据后，关闭通道
	close(ch)
}
