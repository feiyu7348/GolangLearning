// @author:zfy
// @date:2023/8/2 19:49

package main

import "fmt"

type T3 struct {
	n int
}

func getT3() T3 {
	return T3{}
}

func main() {
	t := getT3()
	p := &t.n // <=> p = &(t.n)
	*p = 1
	fmt.Println(t.n)
}
