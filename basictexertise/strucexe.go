// author:zfy  date:2023/3/23

package main

import "fmt"

type person3 struct {
	name string
	age  int
}

func main() {
	p := &person3{}

	p.name = "123"
	p.age = 12
	//p = &person3{
	//	age: 12,
	//}

	fmt.Println(p)
}
