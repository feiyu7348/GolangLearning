// author:zfy  date:2023/3/23

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var p *Person
	fmt.Println(p)
	//p = &Person{}
	//p.name = "zfy"
	//p.age = 10
	//fmt.Println(*p)

	if p == nil {
		p = &Person{age: 1, name: "a"}
	}
	fmt.Println(p)

}
