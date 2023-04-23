// @author:zfy
// @date:2023/4/23 21:43

package main

import "fmt"

var a bool = true

func main() {
	defer func() {
		fmt.Println("1")
	}()
	if a == true {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}

// 2
// 1
