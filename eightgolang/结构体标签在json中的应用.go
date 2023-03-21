// author:zfy  date:2022/6/26

package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  float64  `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "张柏芝"}}
	jsonString, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error:", err)
		return
	}

	fmt.Printf("jsonString = %s\n", jsonString)

	// 解码的过程
	my_movie := Movie{}
	err = json.Unmarshal(jsonString, &my_movie)
	if err != nil {
		fmt.Println("json unmarshal error:", err)
		return
	}

	fmt.Printf("%v\n", my_movie)
}
