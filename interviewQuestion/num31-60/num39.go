// @author:zfy
// @date:2023/6/2 21:51

package main

import "fmt"

func Foo1(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func main() {
	var x *int = nil
	Foo1(x)
}
