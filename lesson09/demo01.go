package main

import (
	"fmt"
	"os"
)

func main() {

	//错误是意料之中的
	file, err := os.Open("aaa.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Name())

}
