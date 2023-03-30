// author:zfy  date:2023/3/27

package main

import "fmt"

func main() {
	numa := make(map[int]int, 1)
	numa[1] = 1
	numa[2] = 2

	v, ok := numa[1]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Printf("没有v: %d", v)
	}

	// 删除key
	delete(numa, 2)

	fmt.Println(numa[2])
}
