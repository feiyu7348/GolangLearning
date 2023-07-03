// @author:zfy
// @date:2023/7/3 23:08

package main

import "fmt"

func main() {
	fmt.Println(^2) // -3
	var a int8 = 3
	var b uint8 = 3
	var c int8 = -3

	fmt.Printf("^%b=%b %d\n", a, ^a, ^a) // ^11=-100 -4
	fmt.Printf("^%b=%b %d\n", b, ^b, ^b) // ^11=11111100 252
	fmt.Printf("^%b=%b %d\n", c, ^c, ^c) // ^-11=10 2
}
