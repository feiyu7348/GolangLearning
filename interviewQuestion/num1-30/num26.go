// @author:zfy
// @date:2023/5/1 9:12

package main

import "fmt"

const (
	a = iota
	b = iota
)
const (
	name = "name"
	c    = iota
	d    = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
