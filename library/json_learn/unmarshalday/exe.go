// @author:zfy
// @date:2023/4/10 16:48

package unmarshalday

import (
	"encoding/json"
	"fmt"
)

type Country struct {
	Address string `json:"address"`
	ID      string `json:"id"`
}

type CountryV1 struct {
	Address string `json:"address"`
	ID      string `json:"id"`
}

type Window struct {
	Width   int        `json:"width"`
	Height  int        `json:"height"`
	Title   *string    `json:"title"`
	Country []*Country `json:"country"`
}

type Mac struct {
	Width   int          `json:"width"`
	Title   string       `json:"title"`
	Height  int          `json:"height"`
	Country []*CountryV1 `json:"country"`
}

func UnmarshalDay(mac *Mac) *Window {
	window := &Window{}

	bytes, _ := json.Marshal(mac)
	//fmt.Println(bytes)
	err := json.Unmarshal(bytes, window)
	if err != nil {
		fmt.Println("JSON decode error!")
		return nil
	}

	return window
}

func JsonDay1() {
	mac := &Mac{
		Width:  1,
		Height: 2,
		Title:  "1",
		//Country: []*CountryV1{{"1", "2"}, {"2", "3"}},
	}

	window := UnmarshalDay(mac)
	fmt.Printf("%+v\n", window)
	fmt.Printf("%+v\n", window.Title)
	fmt.Printf("%+v\n", window.Country)
	for _, v := range window.Country {
		fmt.Println(v)
	}
}
