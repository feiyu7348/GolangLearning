// @author:zfy
// @date:2023/5/22 22:29

package main

import "fmt"

func main() {
	i := 1
	s := []string{"A", "B", "C"}
	i, s[i-1] = 2, "Z"
	fmt.Printf("s: %v \n", s)
}

// s: [Z B C]
