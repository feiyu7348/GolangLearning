// @author:zfy
// @date:2023/7/1 19:56

package main

import "fmt"

const (
	x uint16 = 120
	y
	s = "abc"
	z
)

func main() {
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
}

//    uint16 120
//    string abc
