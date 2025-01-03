// Package docs
// author: zfy  date: 2024/9/29
package docs

import (
	"context"
	"fmt"
	"time"

	"es_study/global"
	"es_study/models"
)

func DocCreate() {
	user := models.UserModel{
		ID:       12,
		UserName: "lisi",
		//Age:       23,
		NickName:  "夜空中最亮的lisi",
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		//Title:     "今天天气很不错",
	}
	indexResponse, err := global.ESClient.Index().Index(user.Index()).BodyJson(user).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", indexResponse)
}
