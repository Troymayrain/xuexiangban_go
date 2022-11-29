package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	//i := rand.Int()
	i := rand.Intn(10)
	fmt.Println(i)

	for i := 0; i < 10; i++ {
		intn := rand.Intn(100)
		fmt.Println(intn)
	}
}
