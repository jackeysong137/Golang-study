package main

import (
	"fmt"
	"time"
)

func main() {

	//获取当前时间 使用time包

	now := time.Now()
	fmt.Println(now) //2024-09-01 15:53:23.475187 +0800 CST m=+0.000741792
	/*
		golang 时间格式化：模版如下：2006-01-02 15:04:05 等同于java的yyyy-MM-dd HH:mm:ss(go 是2006年1月诞生的)
		2006 年
		01 月
		02 日
		03 时 03是12小时制  15表示24小时制
		04 分
		05 秒
	*/
	//2024-09-01 15:53:23
	var tempt = "2006-01-02 15:04:05"
	dateStr := now.Format(tempt)
	fmt.Println(dateStr) //2024-09-01 16:00:39

	//获取当前时时间蹉
	timestamp := time.Now().Unix()
	fmt.Println("时间蹉：", timestamp)

	//时间蹉转化为时间
	dateTime := time.Unix(timestamp, 0)
	fmt.Println("时间：", dateTime

	//转化为时间
	dateStr = "2024-09-01 16:00:39"
	date, _ := time.ParseInLocation(tempt, dateStr, time.Local)
	fmt.Println("date:", date)
	//定时器 1s执行一次
	ticker := time.NewTicker(time.Second)

	n := 0
	for v := range ticker.C {
		if n >= 5 {
			ticker.Stop()
			break
		}
		fmt.Println(v)
		n++
	}
	//睡眠
	time.Sleep(time.Second)
	n = 0
	for {
		time.Sleep(time.Second)

		if n >= 5 {
			break
		}
		fmt.Println(n)
		n++
	}

}
