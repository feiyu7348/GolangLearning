// author:zfy  date:2022/6/24

package main

import "fmt"

func swap(pa *int, pb *int) {
	var tmp int
	tmp = *pa
	*pa = *pb
	*pb = tmp
}

func main() {
	var a int = 10
	var b int = 20
	swap(&a, &b)
	fmt.Println("a = ", a, "b = ", b)
}
