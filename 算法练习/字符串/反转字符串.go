// author:zfy  date:2022/8/5

package main

import "fmt"

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {
	s1 := "hello"
	s := []byte(s1)
	reverseString(s)
	fmt.Println(s)
}
