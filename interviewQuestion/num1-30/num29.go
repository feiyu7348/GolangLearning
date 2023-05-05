// @author:zfy
// @date:2023/5/5 21:43

package main

import (
	"fmt"
	"time"
)

//func main() {
//	v := []int{1, 2, 3}
//	for i := range v {
//		v = append(v, i)
//	}
//	fmt.Println(v)
//}
func main() {

	var m = [...]int{1, 2, 3}
	//
	//for i, v := range m {
	//	go func() {
	//		fmt.Println(i, v)
	//	}()
	//}
	for i, v := range m { // 变量 i、v 在每次循环体中都会被重用，而不是重新声明。
		go func(i, v int) {
			fmt.Println(i, v)
		}(i, v)
	}
	time.Sleep(time.Second * 3)
}

//0 1
//1 2
//2 3
