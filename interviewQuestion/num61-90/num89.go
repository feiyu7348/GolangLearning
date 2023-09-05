// @author:zfy
// @date:2023/9/5 21:47

package main

import "fmt"

func main() {
	v := []int{1, 2, 3}
	for i, n := 0, len(v); i < n; i++ {
		v = append(v, i)
	}
	fmt.Println(v)
}
