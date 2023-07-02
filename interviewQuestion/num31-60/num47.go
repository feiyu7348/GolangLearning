// @author:zfy
// @date:2023/7/2 11:01

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

//1
//1
//2 2
//1
