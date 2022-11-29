package main

import (
	"fmt"
	"time"
)

func main() {

	a := 1

	go func() {
		a = 2
		fmt.Println("goroutine -a ", a)
	}()

	a = 3
	fmt.Println("main -a1 ", a)
	time.Sleep(time.Second * 3)
	fmt.Println("main -a ", a)

}
