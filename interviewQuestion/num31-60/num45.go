// @author:zfy
// @date:2023/6/20 22:14

package main

import "fmt"

func main() {
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)
		i, x := 2, 2
		fmt.Println(i, x)
	}
	fmt.Println(x) // print ?
}
