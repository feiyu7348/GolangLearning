// author:zfy  date:2023/3/23

package main

import (
	"fmt"
	"reflect"
)

type person3 struct {
	name string
	age  int
}

func (p person3) IsEmpty() bool {
	return reflect.DeepEqual(p, person3{})
}

func main() {
	p := &person3{}

	if p.IsEmpty() {
		fmt.Println("p is nil")
	}

	p.name = "0"
	p.age = 12
	//p = &person3{
	//	age: 12,
	//}
	if p.IsEmpty() {
		fmt.Println("p is nil")
	}
	fmt.Println(p)
}
