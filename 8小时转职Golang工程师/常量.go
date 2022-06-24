// author:zfy  date:2022/6/24

package main

import "fmt"

const (
	a = 10 * iota
	b
	c
)

func main() {
	const length int = 10
	fmt.Println("length = ", length)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
}
