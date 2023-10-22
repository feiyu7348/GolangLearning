// @author:zfy
// @date:2023/10/22 20:28

package main

import "fmt"

type foo struct{ Val int }

type bar struct{ Val int }

func main() {
	a := &foo{Val: 5}
	b := &foo{Val: 5}
	c := foo{Val: 5}
	d := bar{Val: 5}
	e := bar{Val: 5}
	f := bar{Val: 5}
	fmt.Print(a == b, c == foo(d), e == f)
}
