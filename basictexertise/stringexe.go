// author:zfy  date:2023/3/23

package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	var s []person
	p := person{
		name: "",
		age:  18,
	}

	s = append(s, p)

	fmt.Println(s)
}
