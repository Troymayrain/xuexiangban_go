package main

import "fmt"

func main() {
	//1、创建map存入用户信息
	user1 := make(map[string]string)
	user1["name"] = "kuangshen"
	user1["age"] = "18"
	user1["sex"] = "男"
	user1["addr"] = "重庆"

	user2 := make(map[string]string)
	user2["name"] = "飞哥"
	user2["age"] = "23"
	user2["sex"] = "男"
	user2["addr"] = "湖南"

	user3 := map[string]string{"name": "小红", "age": "34", "sex": "女", "addr": "北京"}
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)
	//2、将map放入切片
	userDatas := make([]map[string]string, 0, 3)
	userDatas = append(userDatas, user1, user2, user3)
	//userDatas = append(userDatas, user2)
	//userDatas = append(userDatas, user3)
	fmt.Println(userDatas)
	//3、遍历切片获取用户信息
	for _, v := range userDatas {
		fmt.Println("name:", v["name"])
		fmt.Println("age:", v["age"])
		fmt.Println("sex:", v["sex"])
		fmt.Println("addr:", v["addr"])
	}

}
