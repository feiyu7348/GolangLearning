// Package indexes
// author: zfy  date: 2024/9/29
package indexes

import (
	"context"
	"fmt"

	"es_study/global"
)

func DeleteIndex(index string) {
	_, err := global.ESClient.DeleteIndex(index).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(index, "索引删除成功")
}
