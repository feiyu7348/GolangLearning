// Package study
// author: zfy  date: 2024/9/29
package main

import (
	"es_study/core"
	"es_study/docs"
)

func main() {
	core.EsConnect()

	//fmt.Println(global.ESClient)

	//indexs.CreateIndex()

	//docs.DocDelete()
	//docs.DocDeleteBatch()
	//docs.DocCreateBatch()
	//docs.DocFind()
	docs.DocUpdate()

}
