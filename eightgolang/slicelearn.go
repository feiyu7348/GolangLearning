// author:zfy  date:2022/6/25

package main

import "fmt"

func main() {
	//slice1 := []int{1, 2, 3, 4}
	//var slice2 []int
	//slice2 = make([]int, 3)
	//slice3 := make([]int, 4)
	//fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)
	//fmt.Printf("slice2 = %v\n", slice2)
	//fmt.Printf("slice3 = %v\n", slice3)
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1)
	fmt.Printf("len = %d,  cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
}
