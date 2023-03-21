// author:zfy  date:2022/6/24

package main

import "fmt"

func main() {
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)
	var c = "100"
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := "abcd"
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)
}
