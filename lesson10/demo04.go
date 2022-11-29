package main

import (
	"fmt"
	"time"
)

func main() {
	//time1()

	now := time.Now()
	//time转字符串
	format := now.Format("2006-01-02 15:04:05")
	format1 := now.Format("2006-01-02 03:04:05 AM")
	format2 := now.Format("2006/01/02 15:04:05")
	fmt.Println(format)
	fmt.Println(format1)
	fmt.Println(format2)

	//字符串转time
	location, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(location)
	inLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-02-03 13:14:15", location)
	fmt.Println(inLocation)

	//time转时间戳
	//秒
	unix := now.Unix()
	//毫秒
	milli := now.UnixMilli()
	//微妙
	micro := now.UnixMicro()
	//纳秒
	nano := now.UnixNano()
	fmt.Println(unix)
	fmt.Println(milli)
	fmt.Println(micro)
	fmt.Println(nano)

	//时间戳转time
	t := time.Unix(unix, 0)
	t1 := time.UnixMilli(milli)
	unixMicro := time.UnixMicro(micro)
	fmt.Println(t)
	fmt.Println(t1)
	fmt.Println(unixMicro)

}

func time1() {
	now := time.Now()
	fmt.Println(now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
}
