// author:zfy  date:2022/6/24

package main

import "fmt"

const (
	x = iota
	y = iota
	z = iota
)

const v = iota
const (
	h, i, j = iota, iota, iota
)
const (
	a       = iota
	b       = "B"
	c       = iota
	d, e, f = iota, iota, iota
	g       = iota
)

func main() {
	fmt.Println(a, b, c, d, e, f, g, h, i, j, v, x, y, z)
}
