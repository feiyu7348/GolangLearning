// author:zfy  date:2022/6/25

package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human
	level int
}

func main() {
	h := Human{"zhang3", "female"}
	h.Eat()
	h.Walk()
}
