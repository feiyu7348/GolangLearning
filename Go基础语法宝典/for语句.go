// author:zfy  date:2022/6/26

package main

import "fmt"

func main() {
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("Sum is equal to", sum)
}
