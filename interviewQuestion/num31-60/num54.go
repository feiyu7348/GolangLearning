// @author:zfy
// @date:2023/8/1 22:31

package main

import "fmt"

type N int

func (n N) test() {
	fmt.Println(n)
}

func main() {
	var n N = 10
	fmt.Println(n)

	n++
	f1 := N.test
	f1(n)

	n++
	f2 := (*N).test
	f2(&n)

	n++
	N.test(n)

	n++
	(*N).test(&n)
}

//10
//11
//12
//13
//14
