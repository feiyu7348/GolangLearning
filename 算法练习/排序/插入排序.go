// author:zfy  date:2022/7/2

package main

import "fmt"

func Insert(s []int) []int {
	size := len(s)
	for i := 1; i < size; i++ {
		tmp := s[i]
		j := i - 1
		for ; j >= 0 && s[j] > tmp; j-- {
			s[j+1] = s[j]
		}
		s[j+1] = tmp
	}
	return s
}

func main() {
	s := []int{7, 2, 4, 1, 6, 8, 3, 0, 5}
	r := Insert(s)
	fmt.Println(r)
}
