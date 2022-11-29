package main

import "fmt"

func main() {
	var map1 = map[string]int{"GO": 100, "JAVA": 99, "C": 89}
	fmt.Println(map1)

	//for range遍历map
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	//map是引用类型
	map2 := map1
	map2["JAVA"] = 0
	fmt.Println(map1)

}
