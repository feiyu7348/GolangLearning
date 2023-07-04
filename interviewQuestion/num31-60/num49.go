// @author:zfy
// @date:2023/7/4 23:14

package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name"`
}

func main() {
	js := `{
         "name":"seekload"
     }`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(p)
}
