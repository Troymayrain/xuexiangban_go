package main

import "fmt"

func main() {
	//map的定义
	var map1 map[int]string
	map2 := make(map[string]string)
	map2["haha"] = "hehe"
	map2["haha1"] = "hehe1"
	map2["haha2"] = "hehe2"

	map3 := map[string]int{"c": 80, "java": 100, "go": 88}

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)

	fmt.Printf("%T,%p\n", map1, map1)
	fmt.Printf("%T,%p\n", map2, map2)
	fmt.Printf("%T,%p\n", map3, map3)

}
