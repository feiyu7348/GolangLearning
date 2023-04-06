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

func ListElemBelongToOtherList(a []string, elems []string) bool {
	mapA := make(map[string]int)
	for i, v := range a {
		mapA[v] = i
	}

	for _, elem := range elems {
		if _, ok := mapA[elem]; ok {
			continue
		} else {
			return false
		}
	}

	return true
}

func isNil(a []string) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {
	var a []string
	//a := make([]string, 0)
	if len(a) != 0 {
		fmt.Println(a[0])
	}
	fmt.Println(a)
}
