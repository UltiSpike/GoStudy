package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// 打印命令行调用时传给程序的所有参数
	fmt.Println(os.Args)
	// 获取环境变量信息
	//fmt.Println(os.Getenv("PATH"))
	//fmt.Println(os.Setenv("aa", "bb"))
	// grep 查找指定文件里 包含参数的行 并获取其标准输出和标准错误的合并结果
	buf, err := exec.Command("grep", "127.0.0.1", "/etc/hosts").CombinedOutput()

	fmt.Println(string(buf), err)
}
