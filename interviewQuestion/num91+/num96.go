// @author:zfy
// @date:2023/10/25 21:31

package main

import "fmt"

type Point struct{ x, y int }

func main() {
	s := []Point{
		{1, 2},
		{3, 4},
	}
	for _, p := range s {
		p.x, p.y = p.y, p.x
	}
	fmt.Println(s)
}
