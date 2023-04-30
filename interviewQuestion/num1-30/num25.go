// @author:zfy
// @date:2023/4/30 23:32

package main

import "fmt"

type Myint int

func (i Myint) PrintInt() {
	fmt.Println(i)
}

func main() {
	var i Myint = 1
	i.PrintInt()
}
