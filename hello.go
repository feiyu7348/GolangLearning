// author:zfy  date:2022/6/23

package main // 声明 main 包，表明当前是一个可执行程序

import (
	"fmt" // 导入内置 fmt 包
)

func main() { // main函数，是程序执行的入口
	fmt.Println("Hello World!") // 在终端打印 Hello World!
	var m1 map[string]string
	m1["1"] = "1"
	fmt.Println(m1["1"])
}
