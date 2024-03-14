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
			break // 当某一趟序列遍历中元素没有发生交换，则证明该序列已经有序，就不再进行后续的排序。
		}
	}
	return arr
}

func main() {
	s := []int{7, 2, 4, 1, 6, 8, 3, 0, 5}
	r := Bubble2(s)
	fmt.Println(r)
}

func Bubble2(arr []int) []int {
	size := len(arr)
	var flag bool
	for i := 0; i < size-1; i++ {
		flag = false
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}

		if flag != true {
			break
		}
	}

	return arr
}
