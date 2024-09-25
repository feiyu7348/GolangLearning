// Package 即时通信项目
// author: zfy  date: 2024/9/25
package main

import "im/router"

func main() {
	r := router.Router()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
