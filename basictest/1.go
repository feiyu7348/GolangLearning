// author:zfy  date:2023/3/23

package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
	}

	//var a struct{}
	//fmt.Println(a)

	a := person{}
	fmt.Printf("a: %+v\n", a)

	var p person
	fmt.Printf("p: %+v\n", p)

	p.age = 10

	fmt.Printf("p: %+v\n", p)

}
