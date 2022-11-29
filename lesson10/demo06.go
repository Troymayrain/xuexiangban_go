package main

import (
	"fmt"
	"time"
)

func main() {

	//定时器
	//tick := time.Tick(time.Second)
	//for t := range tick {
	//	fmt.Println(t)
	//}

	//for i := 0; i < 10; i++ {
	//	fmt.Println(time.Now())
	//	time.Sleep(time.Second)
	//}

	now := time.Now()
	later := now.Add(time.Hour)
	fmt.Println(now)
	fmt.Println(later)

	sub := later.Sub(now)
	fmt.Println(sub)

	fmt.Println(later.Equal(now))
	fmt.Println(later.After(now))
	fmt.Println(later.Before(now))

}
