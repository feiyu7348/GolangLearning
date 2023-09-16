// @author:zfy
// @date:2023/9/16 22:57

package main

import "fmt"

type User1 struct {
	Name string
}

func (u *User1) SetName(name string) {
	u.Name = name
	fmt.Println(u.Name)
}

type Employee User

func main() {
	employee := new(Employee)
	employee.SetName("Jack")
}
