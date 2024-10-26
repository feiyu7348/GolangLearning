// Package template_learn
// author: zfy  date: 2024/10/26
package main

import (
	"os"
	"text/template"
)

func main() {
	// 构建模板
	tmpl, err := template.New("").Parse("你好，{{ . }}")
	if err != nil {
		panic(err)
	}

	// 输出文本, 最终输出 "你好，冰镇"
	err = tmpl.Execute(os.Stdout, "冰镇")
	if err != nil {
		panic(err)
	}
}
