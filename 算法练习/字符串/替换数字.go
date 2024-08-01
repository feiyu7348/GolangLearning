// @author:zfy
// @date:2024/8/1 19:35

package main

import "fmt"

func replaceNumber(s string) string {
	size := len(s)
	insertElement := []byte{'n', 'u', 'm', 'b', 'e', 'r'}
	replaceStringed := make([]byte, 0)
	sBytes := []byte(s)

	for i := 0; i < size; i++ {
		if sBytes[i] <= '9' && sBytes[i] >= '0' {
			replaceStringed = append(replaceStringed, insertElement...)
		} else {
			replaceStringed = append(replaceStringed, sBytes[i])
		}
	}

	return string(replaceStringed)
}

func main() {
	fmt.Println(replaceNumber("129r"))
}
