// @author:zfy
// @date:2023/7/28 19:19

package main

import "fmt"

type T1 struct {
	n int
}

func main() {
	m := make(map[int]T1)
	t := T1{1}
	m[0] = t
	fmt.Println(m[0].n)
}
