package main

import "fmt"

func main() {
	//通过数组创建切片
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//此时切片s1的底层是当前数组，【start包含，end不包含】
	s1 := arr[:5]
	s2 := arr[3:5]
	s3 := arr[5:]
	//通过数组创建切片，长度是：【start,end】，容量是：【start,最后】
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("len:%d,cap:%d\n", len(s2), cap(s2))
	fmt.Printf("len:%d,cap:%d\n", len(s3), cap(s3))

	//修改数组的值
	arr[0] = 100
	fmt.Println(arr, s1)

	//修改切片的值
	s1[0] = 200
	fmt.Println(arr, s1)

	fmt.Printf("s1:%p,arr:%p\n", s1, &arr)
	fmt.Printf("s2:%p,arr:%p\n", s2, &arr)

	//切片一旦发生扩容，就会指向新的数组，此时再改变切片不会影响原数组
	s1 = append(s1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	fmt.Println(arr, s1)

	s1[0] = 300
	fmt.Println(arr, s1)

}
