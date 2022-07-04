// author:zfy  date:2022/7/4

package main

import "fmt"

func removeElements(nums []int, val int) int {
	length := len(nums)
	res := 0
	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 2}
	val := 2
	fmt.Println(removeElements(nums, val))
}
