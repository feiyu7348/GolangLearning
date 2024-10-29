// Package indexes
// author: zfy  date: 2024/9/29
package indexes

import (
	"context"
	"es_study/global"
	"es_study/models"
	"fmt"
)

func CreateIndex() {
	index := "user_index"
	if ExistsIndex(index) {
		// 索引存在，先删除，在创建
		DeleteIndex(index)
	}

	createIndex, err := global.ESClient.
		CreateIndex(index).
		BodyString(models.UserModel{}.Mapping()).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(createIndex.Index, "索引创建成功")
}
