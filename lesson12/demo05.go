package main

import (
	"fmt"
	"sync"
	"time"
)

//不要以共享内存的方式通信，而要以通信的方式共享内存

// 创建锁
var mutex sync.Mutex

// 创建等待组
var wg sync.WaitGroup

// 定义一个全局变量，总票数
var ticket int = 10

// 售票问题
func main() {

	wg.Add(4)

	go saleTickets("售票窗口1")
	go saleTickets("售票窗口2")
	go saleTickets("售票窗口3")
	go saleTickets("售票窗口4")

	//time.Sleep(time.Second * 5)

	wg.Wait()
	fmt.Println("main goroutine执行完！")
}

// 售票函数
func saleTickets(name string) {
	defer wg.Done()
	for {
		//检查之前先上锁
		mutex.Lock()
		if ticket > 0 {
			fmt.Println(name+"剩余票数：", ticket)
			ticket--
			time.Sleep(time.Millisecond * 500)
		} else {
			//操作完后释放锁
			mutex.Unlock()
			fmt.Println("票已经卖完!")
			break
		}
		//操作完后释放锁
		mutex.Unlock()
	}
}
