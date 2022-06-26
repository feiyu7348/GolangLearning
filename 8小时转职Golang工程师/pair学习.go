// author:zfy  date:2022/6/26

package main

import "fmt"

func main() {
	var a string
	a = "abcd"
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}
