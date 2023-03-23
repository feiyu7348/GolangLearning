// author:zfy  date:2023/3/23

package main

import (
	"encoding/json"
	"fmt"
)

type person1 struct {
	name string
	age  int
}

func main() {
	p := &person1{
		name: "1",
		age:  1,
	}

	var v []byte
	//var a []person1

	err := json.Unmarshal(v, p)
	if err != nil {
		fmt.Println(v)
		fmt.Println(p)
	}

}
