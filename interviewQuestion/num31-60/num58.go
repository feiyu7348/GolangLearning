// @author:zfy
// @date:2023/8/5 22:40

package main

import "fmt"

type T4 struct {
	x int
	y *int
}

func main() {

	i := 20
	t := T4{10, &i}

	p := &t.x

	*p++
	*p--

	t.y = p

	fmt.Println(*t.y)
}
