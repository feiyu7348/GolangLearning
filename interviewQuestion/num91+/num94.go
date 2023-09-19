// @author:zfy
// @date:2023/9/19 21:58

package main

import "fmt"

var nil = new(int)

func main() {
	var p *int
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}
