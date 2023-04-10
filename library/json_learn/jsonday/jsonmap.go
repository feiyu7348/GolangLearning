// author:zfy  date:2023/3/23

package jsonday

import (
	"encoding/json"
	"fmt"
)

func MapToJson() {
	fileCount := map[string]int{
		"cpp": 10,
		"js":  8,
		"go":  10,
	}
	bytes, _ := json.Marshal(fileCount)
	fmt.Println(string(bytes))
}
