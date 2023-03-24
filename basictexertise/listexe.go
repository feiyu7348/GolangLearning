// author:zfy  date:2023/3/23

package main

import (
	"container/list"
	"fmt"
)

type person2 struct {
	name string
	age  int
}

func main() {
	t := list.New()
	//var s []string
	var p []person2

	fmt.Println(t)
	//fmt.Println(s)
	fmt.Println(p)

	//定义变量
	var s []string
	fmt.Printf("1:nil=%t, len=%d, cap=%d\n", s == nil, len(s), cap(s))
	fmt.Println(s)

	//组合字面量方式
	s = []string{}
	fmt.Printf("2:nil=%t, len=%d, cap=%d\n", s == nil, len(s), cap(s))
	fmt.Println(s)

	//make方式
	s = make([]string, 0)
	fmt.Printf("3: nil=%t, len=%d, cap=%d\n", s == nil, len(s), cap(s))
	fmt.Println(s)

	s = new([{}, {]])
	fmt.Println(s)

}
