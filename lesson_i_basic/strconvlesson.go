package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 按指定进制转化浮点数和整数
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	n, _ := strconv.ParseInt("111", 10, 64)
	fmt.Println(n)

	n2, _ := strconv.Atoi("123")
	fmt.Println(n2)
	n2, err := strconv.Atoi("AAA")
	fmt.Println(n2, err)
}
