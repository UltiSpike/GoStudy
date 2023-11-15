package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// 1.生成一个随机数
// 2.接收输入
// 3.如果》
// 4.如果《

func main() {

	n := rand.Intn(100)

	var cnt int
	reader := bufio.NewReader(os.Stdin)
	for {
		cnt += 1
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error:", err)
			return
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input:", err)
			return
		}
		if guess > n {
			fmt.Println("smaller ")
		} else if guess < n {
			fmt.Println("bigger")
		} else {
			fmt.Printf("Congratulation! The number is %d , conut: %d", n, cnt)
			break
		}
	}

}
