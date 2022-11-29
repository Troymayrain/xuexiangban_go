package main

import (
	"fmt"
	"strings"
)

// strings包的使用
func main() {
	str := "xuexiangban,kuangsheng"

	//Contains是否包含指定内容
	fmt.Println(strings.Contains(str, "k"))
	fmt.Println(strings.Contains(str, "z"))
	fmt.Println(strings.Contains(str, "ka"))
	//ContainsAny是否包含指定的任意一个内容
	fmt.Println(strings.ContainsAny(str, "ka"))
	//Count统计指定内容在字符串中出现的次数
	fmt.Println(strings.Count(str, "a"))
	fmt.Println(strings.Count(str, "sheng"))

	//HasPrefix判断字符串以xx开头
	fileName := "20221127.mp4"
	if strings.HasPrefix(fileName, "2023") {
		fmt.Println(fileName)
	}
	//HasSuffix判断字符串以xx结尾
	if strings.HasSuffix(fileName, ".mp4") {
		fmt.Println("文件是mp4格式的")
	}

	//Index查找指定字符串第一次出现的位置
	fmt.Println(strings.Index(str, "n"))
	fmt.Println(strings.Index(str, "z"))
	//IndexAny查找指定任意字符串第一次出现的位置
	fmt.Println(strings.IndexAny(str, "nz"))
	fmt.Println(strings.IndexAny(str, "zz"))
	//LastIndex查找指定字符串最后一次出现的位置
	fmt.Println(strings.LastIndex(str, "n"))
	fmt.Println(strings.LastIndex(str, "zz"))

	//Join字符串拼接
	s := []string{"a", "b", "c", "d", "e"}
	str1 := strings.Join(s, "-")
	fmt.Println(str1)

	//Split字符串拆分
	s1 := strings.Split(str1, "-")
	fmt.Println(s1)

	//字符串重复
	repeat := strings.Repeat("a", 5)
	fmt.Println(repeat)

	//字符串替换
	replace := strings.Replace(str, "u", "*", 1)
	fmt.Println(replace)

	//大小写转换
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))

	//字符串截取
	str3 := str[0:5]
	str4 := str[5:]
	fmt.Println(str3)
	fmt.Println(str4)

}
