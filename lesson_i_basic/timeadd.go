package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Println(now)
	// 设置一个时间 时区
	t := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC)
	t2 := time.Date(2022, 3, 27, 2, 25, 36, 0, time.UTC)
	// 不同的Time.time时间对象
	// 默认UTC
	fmt.Println(t)
	// 按年-月-天-时-秒输出
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	// 按自己选择的格式输出
	// 格式方法1 "2006-01-02 15:04:05"
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	diff := t2.Sub(t)
	// 计算时差 默认最大
	fmt.Println(diff)
	// 显示分钟差距 秒钟差距
	fmt.Println(diff.Minutes(), diff.Seconds())
	// 将string 转为指定格式的时间对象
	t3, _ := time.Parse("2006-01-02 15:04:05", "2022-03-27 01:25:36")
	// t3 是一个时间对象
	fmt.Println(t3 == t)
	// Unix 时间戳表示从 1970 年 1 月 1 日 00:00:00 UTC 到给定时间的秒数。
	fmt.Println(now.Unix())
}
