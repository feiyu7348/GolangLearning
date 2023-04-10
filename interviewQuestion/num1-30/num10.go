// @author:zfy
// @date:2023/4/10 9:20

package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[3:4:4]
	fmt.Println(b[0])
}
