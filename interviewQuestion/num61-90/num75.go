// @author:zfy
// @date:2023/8/21 22:31

package main

import "fmt"

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover:%#v", r)
		}
	}()
	panic(1)
	panic(2)
}

func main() {
	f()
}
