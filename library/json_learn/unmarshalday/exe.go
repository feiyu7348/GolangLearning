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

type Window struct {
	Width   uint64     `json:"width"`
	Height  uint64     `json:"height"`
	Title   []string   `json:"title"`
	country []*Country `json:"country"`
}

type Mac struct {
	Width   int        `json:"width"`
	Height  int        `json:"height"`
	Title   []string   `json:"title"`
	country []*Country `json:"country"`
}

func UnmarshalDay(mac *Mac) *Window {
	var window *Window

	bytes, _ := json.Marshal(mac)
	err := json.Unmarshal(bytes, &window)
	if err != nil {
		fmt.Println("JSON decode error!")
		return nil
	}

	return window
}

func JsonDay1() {
	mac := &Mac{
		Width:   1,
		Height:  2,
		Title:   []string{"3", "4"},
		country: []*Country{{"1", "2"}, {"2", "3"}},
	}

	window := UnmarshalDay(mac)
	fmt.Printf("%+v", window)
}
