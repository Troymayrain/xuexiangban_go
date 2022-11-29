package main

import (
	"fmt"
	"sync"
)

// 创建等待组
var wg1 sync.WaitGroup

func main() {

	wg1.Add(2)

	go test1()
	go test2()

	fmt.Println("main ing")
	wg1.Wait()
	fmt.Println("WaitGroup 解除")

}

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1-", i)
	}
	wg1.Done()
}
func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2-", i)
	}
	wg1.Done()
}
