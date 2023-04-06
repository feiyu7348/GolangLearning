// @author:zfy
// @date:2023/4/6 9:53

package main

import "fmt"

type MyInt1 int   // 基于类型 int 创建了新类型 MyInt1
type MyInt2 = int // 创建了 int 的类型别名 MyInt2

func main() {
	var i int = 0

	// 将 int 类型的变量赋值给 MyInt1 类型的变量，Go 是强类型语言，编译当然不通过
	//var i1 MyInt1 = i
	// 赋值可以使用强制类型转化
	var i1 MyInt1 = MyInt1(i)

	// MyInt2 只是 int 的别名，本质上还是 int，可以赋值
	var i2 MyInt2 = i

	fmt.Println(i1)
	fmt.Println(i2)
}
