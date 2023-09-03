// @author:zfy
// @date:2023/9/3 21:12

package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["foo"]++
	fmt.Println(m["foo"])
}
