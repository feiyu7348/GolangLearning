// @author:zfy
// @date:2023/3/31 17:54

package main

import "fmt"

// 1.
func a() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

// 2.
func b() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}

func main() {
	a() // [0 0 0 0 0 1 2 3]
	b() // [1 2 3 4]
}
