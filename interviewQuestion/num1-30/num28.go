// @author:zfy
// @date:2023/5/5 21:48

package main

import "fmt"

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Println(*p)
}

//func main() { 报错
//	p, err := foo()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	bar()
//	fmt.Println(*p)
//}
func main() {
	var err error
	p, err = foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}
