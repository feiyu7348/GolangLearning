// author:zfy  date:2022/6/26

package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}
