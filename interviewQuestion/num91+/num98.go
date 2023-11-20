// @author:zfy
// @date:2023/11/20 21:45

package main

import "fmt"

func main() {
	a := 1
	for i := 0; i < 5; i++ {
		a := a + 1
		a = a * 2
	}
	fmt.Println(a)
}
