// author:zfy  date:2023/3/24

package jsonday

import (
	"encoding/json"
	"fmt"
)

type Seller struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
}

type product struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Seller Seller `json:"seller"`
	Price  int    `json:"price"`
}

func ListToJson() {
	products := []product{
		product{
			Id:     50,
			Name:   "Writing Book",
			Seller: Seller{1, "ABC Company", "US"},
			Price:  100,
		},
		product{
			Id:     51,
			Name:   "Kettle",
			Seller: Seller{20, "John Store", "DE"},
			Price:  500,
		},
	}
	bytes, _ := json.Marshal(products)
	fmt.Println(string(bytes))
}
