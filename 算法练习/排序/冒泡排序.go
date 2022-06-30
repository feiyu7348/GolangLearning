// author:zfy  date:2022/6/30

package main

import "fmt"

func Bubble(arr []int) []int {
	size := len(arr)
	var swapped bool
	for i := 0; i < size-1; i++ {
		swapped = false
		for j := 0; j < size-1-i; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if swapped != true {
			break
		}
	}
	return arr
}

func main() {
	s := []int{7, 2, 4, 1, 6, 8, 3, 0, 5}
	r := Bubble(s)
	fmt.Println(r)
}
