// author:zfy  date:2022/6/29

package main

import "fmt"

func liner_search(s []int, v int) int {
	for i := 0; i < len(s); i++ {
		if s[i] == v {
			return i
		}
	}
	return -1
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	v := 3
	r := liner_search(s, v)
	fmt.Println(r)
}
