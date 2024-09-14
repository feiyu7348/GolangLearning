// author:zfy  date:2022/6/29

package main

import "fmt"

func binarySearch(arr []int, target int, l int, r int) int {
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	index := binarySearch(arr, 4, 0, len(arr)-1)
	fmt.Println(index)
	fmt.Println(search(arr, 4))
}
