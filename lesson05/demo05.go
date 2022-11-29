package main

import "fmt"

func main() {
	//数组冒泡排序
	arr := [6]int{8, 1, 23, 4, 6, 0}
	sort(arr, "desc")

}

func sort(arr [6]int, c string) {
	for j := 1; j <= len(arr); j++ {
		for i := 0; i < len(arr)-j; i++ {
			if c == "asc" {
				if arr[i] > arr[i+1] {
					arr[i], arr[i+1] = arr[i+1], arr[i]
				}
			}
			if c == "desc" {
				if arr[i] < arr[i+1] {
					arr[i], arr[i+1] = arr[i+1], arr[i]
				}
			}
		}
		fmt.Println(arr)
	}
	fmt.Println("排序完的数组：", arr)
}
