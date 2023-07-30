// @author:zfy
// @date:2023/7/30 18:25

package main

import "fmt"

type T2 struct {
	n int
}

func (t *T2) Set(n int) {
	t.n = n
}

func getT() T2 {
	return T2{}
}

func main() {
	t := getT()
	t.Set(2)
	fmt.Println(t.n)
}
