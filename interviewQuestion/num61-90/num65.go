// @author:zfy
// @date:2023/8/11 20:49

package main

import "fmt"

func main() {
	defer func() {
		fmt.Print(recover())
	}()
	defer func() {
		defer func() {
			fmt.Print(recover())
		}()
		panic(1)
	}()
	defer recover()
	panic(2)
}
