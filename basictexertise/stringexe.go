// author:zfy  date:2023/3/23

package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

var a = "titleNow"

func main() {
	//var s []person
	//p := person{
	//	name: "",
	//	age:  18,
	//}
	//
	//s = append(s, p)

	fmt.Println(strings.ToUpper(a))
	fmt.Println(strings.ToUpper(a[:1]) + a[1:])
	s := []string{"1", "2", "3"}
	var builder strings.Builder
	for _, v := range s {
		builder.WriteString(v + "\n")
	}
	fmt.Println(builder.String())
}
