// @author:zfy
// @date:2023/8/29 22:55

package main

import "fmt"

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func main() {
	d1 := data{"one"}
	d1.print()

	var in printer = data{"two"}
	in.print()
}
