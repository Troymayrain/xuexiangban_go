package main

import "fmt"

func main() {
	var map1 map[int]string //nil
	map1 = make(map[int]string)
	map1[1] = "xuexiangban"
	map1[3] = "kuangsheng"
	map1[5] = "feige"
	fmt.Println(map1)

	fmt.Println(map1[1])
	fmt.Println(map1[2])

	//判断map中的key是否存在,ok-idom
	value, ok := map1[5]
	if ok {
		fmt.Println("map key 存在，value:", value)
	} else {
		fmt.Println("map key 不存在")
	}

	delete(map1, 1)
	fmt.Println(map1)

	fmt.Println(len(map1))

}
