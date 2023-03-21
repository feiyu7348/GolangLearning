// author:zfy  date:2022/6/26

// 断言
package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is caller...")
	fmt.Println(arg)

	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is a string type, value = ", value)
		fmt.Println("value type is %T\n", value)
	}
}

type Book1 struct {
	author string
}

func main() {
	book := Book1{"Golang"}
	myFunc(book)
}
