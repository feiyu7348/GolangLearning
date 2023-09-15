// @author:zfy
// @date:2023/9/15 19:44

package main

import "fmt"

func min(a int, b uint) {
	var min = 0
	min = copy(make([]struct{}, a), make([]struct{}, b))
	fmt.Printf("The min of %d and %d is %d\n", a, b, min)
}

func main() {
	min(1225, 256)
}
