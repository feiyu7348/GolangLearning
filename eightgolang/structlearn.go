// author:zfy  date:2022/6/25

package main

import "fmt"

type myint int
type Book struct {
	title  string
	author string
}

func main() {
	var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	var bool1 Book
	bool1.title = "Go"
	bool1.author = "zhang"
	fmt.Printf("%v\n", bool1)
}
