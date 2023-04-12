// author:zfy  date:2023/3/24

package jsonday

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Id     string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func StructToJson() {
	myBook := Book{Id: "Golang", Author: "zfy", Year: 2023}
	bytes, _ := json.Marshal(myBook)
	fmt.Println(string(bytes))
}
