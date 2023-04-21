// @author:zfy
// @date:2023/4/21 22:02

package main

import "fmt"

type Person struct {
	age int
}

func main() {
	person := &Person{28}

	// 1.
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age)
	}()

	person = &Person{29}
}

// 29 28 28
