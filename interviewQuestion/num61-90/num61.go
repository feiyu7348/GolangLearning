// @author:zfy
// @date:2023/8/7 22:18

package main

import "fmt"

func main() {
	var i = 1
	var s = []int{1, 2}
	i, s[i] = 0, 3
	fmt.Println(s[0] + s[1])

	var k = 9
	for k = range []int{} {
	}
	fmt.Println(k)

	for k = 0; k < 3; k++ {
	}
	fmt.Println(k)

	for k = range (*[3]int)(nil) {
	}
	fmt.Println(k)
}
