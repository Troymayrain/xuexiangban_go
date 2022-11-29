package main

import "fmt"

//程序-程序是静态的，死的
//进程-程序跑起来产生进程
//线程-线程是轻量级的进程，所有线程共享同一进程的资源
//协程（goroutine）-协程是轻量级的线程

//串行：挨个执行
//并行：同时执行，一般多核cpu
//并发：多个goroutine快速交替执行

func main() {
	//开启一个goroutine
	go hello()

	for i := 0; i < 100; i++ {
		fmt.Println("main----", i)
	}
}

func hello() {
	for i := 0; i < 100; i++ {
		fmt.Println("hello----", i)
	}
}
