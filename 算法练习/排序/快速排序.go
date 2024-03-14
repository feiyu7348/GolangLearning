// author:zfy  date:2022/7/2

package main

import "fmt"

func partition(s []int, left int, right int) int {
	tmp := s[left]
	for left < right {
		for left < right && s[right] >= tmp {
			right--
		}
		s[left] = s[right]
		for left < right && s[left] <= tmp {
			left++
		}
		s[right] = s[left]
	}
	s[left] = tmp
	return left
}

func Quick(s []int, left int, right int) []int {
	if left < right {
		mid := partition(s, left, right)
		Quick(s, left, mid-1)
		Quick(s, mid+1, right)
	}
	return s
}

func main() {
	s := []int{7, 2, 4, 1, 6, 8, 3, 0, 5}
	right := len(s) - 1
	r := Quick(s, 0, right)
	fmt.Println(r)
}

func quickSort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	pivot := s[0]
	left, right := make([]int, 0), make([]int, 0)
	for _, ele := range s[1:] {
		if ele <= pivot {
			left = append(left, ele)
		} else {
			right = append(right, ele)
		}
	}

	return append(quickSort(left),
		append([]int{pivot}, quickSort(right)...)...)
}
