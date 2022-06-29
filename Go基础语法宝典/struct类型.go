// author:zfy  date:2022/6/26

package main

import "fmt"

type person struct {
	name string
	age  int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var P person
	P.name = "fei"
	P.age = 26
	fmt.Printf("The person'name is %s\n", P.name)

	var tom person
	tom.name, tom.age = "Tom", 18
	bob := person{age: 25, name: "Bob"}
	pual := person{"Paul", 26}
	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, pual)
	bp_Older, bp_diff := Older(bob, pual)
	fmt.Printf("of %s and %s, %s is older by %d year\n",
		tom.name, bob.name, tb_Older, tb_diff)
	fmt.Printf("of %s and %s, %s is older by %d year\n",
		tom.name, pual.name, tp_Older, tp_diff)
	fmt.Printf("of %s and %s, %s is older by %d year\n",
		bob.name, pual.name, bp_Older, bp_diff)
}
