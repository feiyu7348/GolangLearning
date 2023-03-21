// author:zfy  date:2022/6/24

package main

import "fmt"

func main() {
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main:hello go 1")
	fmt.Println("main:hello go 2")
}
