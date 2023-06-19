// @author:zfy
// @date:2023/6/19 21:52

package main

import "fmt"

type info struct {
	result int
}

func work() (int, error) {
	return 13, nil
}
func main() {
	var data info
	var err error
	data.result, err = work() //ok
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}
