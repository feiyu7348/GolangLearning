// @author:zfy
// @date:2023/8/3 22:40

package main

import "fmt"

func main() {
	s := make([]int, 3, 9)
	fmt.Println(len(s))
	s2 := s[4:8]
	fmt.Println(len(s2))
}
