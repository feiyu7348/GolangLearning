// @author:zfy
// @date:2023/4/11 9:15

package main

import "fmt"

func main() {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
}

// nil
// 0
