// @author:zfy
// @date:2023/4/9 14:42

package main

import "fmt"

func hello1(num ...int) {
	num[0] = 18
}

func main() {
	i := []int{5, 6, 7}
	hello1(i...)
	fmt.Println(i[0])
}
