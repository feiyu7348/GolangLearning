// author:zfy  date:2023/3/23

package main

import "fmt"

func removeItem(a []int, elem int) []int {
	j := 0
	for _, v := range a {
		if v != elem {
			a[j] = v
			j++
		}
	}

	return a[:j]
}

func removeItem1(a []int, elem int) []int {
	var aNew []int
	for _, v := range a {
		if v != elem {
			aNew = append(aNew, v)
		}
	}

	return aNew
}

func main() {
	a := []int{1, 2, 3, 4}
	elem := 2
	fmt.Println(removeItem(a, elem))
	b := []int{1, 2, 3, 4}
	elemb := 2
	fmt.Println(removeItem1(b, elemb))
}
