// author:zfy  date:2022/6/26

package main

import "fmt"

func main() {
	//修改字符串内容
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	s = "C" + s[1:]
	//反引号括起来的字符串为 Raw 字符串
	e := `hello
			world`
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", s2)
	fmt.Printf("%s\n", e)

	//使用加号连接字符串
	a := "hello,"
	b := " world"
	d := a + b
	fmt.Printf("%s\n", d)
}
