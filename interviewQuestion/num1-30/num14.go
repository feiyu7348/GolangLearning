// @author:zfy
// @date:2023/4/14 21:44

package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	p := 1
	incr(&p)
	fmt.Println(p)
}
