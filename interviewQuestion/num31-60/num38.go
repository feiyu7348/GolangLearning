// @author:zfy
// @date:2023/5/31 23:11

package main

import "fmt"

//func main() {
//	x := []string{"a", "b", "c"}
//	for v := range x {
//		fmt.Print(v) // 012
//	}
//	y := []string{"a", "b", "c"}
//	for _, v := range y {
//		fmt.Print(v) //输出 abc
//	}
//}

type User struct{}
type User1 User   // 基于类型 User 创建了新类型 User1
type User2 = User // 创建了 User 的类型别名 User2

func (i User1) m1() {
	fmt.Println("m1")
}
func (i User) m2() {
	fmt.Println("m2")
}

func main() {
	var i1 User1
	var i2 User2
	i1.m1()
	i2.m2()
}
