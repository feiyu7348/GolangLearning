// author:zfy  date:2023/3/30

package main

import "fmt"

// NUm1 数组使用值拷贝传参
func NUm1() {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr) // [7 2 3]
	}(x)
	fmt.Println(x) // [1 2 3]    // 并不是你以为的 [7 2 3]
}

// Num2 传址会修改原数据
func Num2() {
	x := [3]int{1, 2, 3}

	func(arr *[3]int) {
		(*arr)[0] = 7
		fmt.Println(arr) // &[7 2 3]
	}(&x)
	fmt.Println(x) // [7 2 3]
}

// Num3 会修改 slice 的底层 array，从而修改 slice
func Num3() {
	x := []int{1, 2, 3}
	func(arr []int) {
		arr[0] = 7
		fmt.Println(x) // [7 2 3]
	}(x)
	fmt.Println(x) // [7 2 3]
}

func main() {
	NUm1()
	Num2()
	Num3()
}
