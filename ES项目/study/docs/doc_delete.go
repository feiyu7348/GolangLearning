// Package docs
// author: zfy  date: 2024/9/29
package docs

import (
	"context"
	"fmt"

	"es_study/global"
	"es_study/models"
)

func DocDelete() {
	deleteResponse, err := global.ESClient.Delete().
		Index(models.UserModel{}.Index()).Id("tmcqfYkBWS69Op6Q4Z0t").Refresh("true").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(deleteResponse)
}
