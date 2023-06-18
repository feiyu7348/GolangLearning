// @author:zfy
// @date:2023/6/18 9:48

package main

func main() {
	x := interface{}(nil)
	y := (*int)(nil)
	a := y == x
	b := y == nil
	_, c := x.(interface{})
	println(a, b, c)
}

// false true false
