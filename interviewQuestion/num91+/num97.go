// @author:zfy
// @date:2023/11/19 15:11

package main

import "fmt"

type Foo struct {
	val int
}

func (f Foo) Inc(inc int) {
	f.val += inc
}

func main() {
	var f Foo
	f.Inc(100)
	fmt.Println(f.val)
}
