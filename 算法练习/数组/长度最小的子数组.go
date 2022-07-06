// author:zfy  date:2022/7/6

package main

import "fmt"

func minSubArraylength(target int, nums []int) int {
	i := 0
	l := len(nums)
	sum := 0
	result := l + 1
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArraylength(target, nums))
}
