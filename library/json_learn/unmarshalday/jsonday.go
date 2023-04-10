// @author:zfy
// @date:2023/4/10 16:35

package unmarshalday

import (
	"encoding/json"
	"fmt"
)

func JsonDay() {
	jsonInput := `{
        "width": 500,
        "height": 200,
        "title": "Hello Go!"
    }`

	var window Window
	err := json.Unmarshal([]byte(jsonInput), &window)
	if err != nil {
		fmt.Println("JSON decode error!")
		return
	}

	fmt.Println(window)       // {500 200 Hello Go!}
	fmt.Printf("%+v", window) // {Width:500 Height:200 Title:Hello Go!}
}
