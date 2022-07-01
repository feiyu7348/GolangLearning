// author:zfy  date:2022/7/1

package main

import "fmt"

func Select(s []int) []int {
	size := len(s)
	for i := 0; i < size-1; i++ {
		min_index := i
		for j := i + 1; j < size; j++ {
			if s[j] < s[min_index] {
				min_index = j
			}
		}
		s[i], s[min_index] = s[min_index], s[i]
	}
	return s
}

func main() {
	s := []int{7, 2, 4, 1, 6, 8, 3, 0, 5}
	r := Select(s)
	fmt.Println(r)
}
