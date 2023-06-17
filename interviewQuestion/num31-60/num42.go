// @author:zfy
// @date:2023/6/7 21:41

package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)

	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}

		r = append(r, v)
	}

	fmt.Println(r)
}

// [1 2 3 4 5]
// or range 时会使用 a 的副本 a' 参与循环，副本的 len 依旧是 5
