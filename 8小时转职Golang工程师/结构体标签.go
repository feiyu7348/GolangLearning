// author:zfy  date:2022/6/26

package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `json:"name" doc:"我的名字"`
	Sex  string `json:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tapstring := t.Field(i).Tag.Get("json")
		fmt.Println("json", tapstring)
	}
}

func main() {
	var re resume
	findTag(&re)
}
