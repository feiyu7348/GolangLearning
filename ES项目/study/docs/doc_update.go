// Package docs
// author: zfy  date: 2024/9/29
package docs

import (
	"es_study/global"
	"es_study/models"
	"fmt"
)

func DocUpdate() {
	res, err := global.ESClient.Update().Index(models.UserModel{}.Index()).
		Id("vmdnfYkBWS69Op6QEp2Y").Doc(map[string]any{
		"user_name": "你好枫枫",
	}).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", res)
}
