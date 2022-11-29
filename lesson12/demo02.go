package main

import (
	"fmt"
	"runtime"
)

// runtime包的使用
func main() {

	fmt.Println("GOROOT:", runtime.GOROOT())
	fmt.Println("操作系统:", runtime.GOOS)
	fmt.Println("CPU个数", runtime.NumCPU())

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("goroutine")
		}
	}()

	for i := 0; i < 100; i++ {
		//让出时间片，让别的goroutine先执行，但不一定能成功
		runtime.Gosched()
		fmt.Println("main")
	}

}
