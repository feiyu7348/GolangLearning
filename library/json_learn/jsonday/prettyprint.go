// author:zfy  date:2023/3/24

package jsonday

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint() {
	book := product{
		Id:     50,
		Name:   "Writing Book",
		Seller: Seller{1, "ABC Company", "US"},
		Price:  100,
	}
	bytes, _ := json.MarshalIndent(book, "", "\t")
	fmt.Println(string(bytes))
}
