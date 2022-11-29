package main

import "fmt"

func main() {
	//定义数组
	var nums [4]int

	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4

	fmt.Printf("%T\n", nums)
	fmt.Println(nums[1])
	fmt.Println(nums[2])

	//数组长度
	fmt.Println(len(nums))
	//数组容量
	fmt.Println(cap(nums))

	fmt.Println(nums)
	nums[0] = 100
	fmt.Println(nums)

}
