// @author:zfy
// @date:2023/8/6 10:39

package main

import "fmt"

type N1 int

func (n *N1) test() {
	fmt.Println(*n)
}

func main() {
	var n N1 = 10
	p := &n

	n++
	f1 := n.test

	n++
	f2 := p.test

	n++
	fmt.Println(n)

	f1()
	f2()
}
