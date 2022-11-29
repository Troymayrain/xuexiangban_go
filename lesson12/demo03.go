package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go func() {
		fmt.Println("start")
		test()
		fmt.Println("end")
	}()

	time.Sleep(time.Second * 2)

}

func test() {
	defer fmt.Println("test defer")
	//return只会终止函数
	//return
	//Goexit会终止goroutine
	runtime.Goexit()
	fmt.Println("test")
}
