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
	//a := []int{1, 2, 3, 4}
	//elem := 2
	//fmt.Println(removeItem(a, elem))
	//b := []int{1, 2, 3, 4}
	//elemb := 2
	//fmt.Println(removeItem1(b, elemb))

	var arrayA []int
	fmt.Printf("arrayA: %d\n", arrayA)
	if arrayA == nil {
		fmt.Println("arrayA is nil")
	}

	arrayB := make([]int, 0)
	fmt.Printf("arrayB: %d\n", arrayB)
	if arrayB == nil {
		fmt.Println("arrayB is nil")
	}

	arrayC := []int{}
	fmt.Printf("arrayC: %d\n", arrayC)
	if arrayC == nil {
		fmt.Println("arrayC is nil")
	}

	a := 1
	arrayA = append(arrayA, a)
	fmt.Printf("arrayA: %d\n", arrayA)

	b := 2
	arrayB = append(arrayB, b)
	fmt.Printf("arrayB: %d\n", arrayB)

	c := 3
	d := 4
	arrayC = append(arrayC, c, d)
	fmt.Printf("arrayC: %d\n", arrayC)
}
