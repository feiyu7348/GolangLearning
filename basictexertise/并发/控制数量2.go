// Package main
// author: zfy  date: 2024/9/26
package main

import (
	"log"
	"time"

	"github.com/Jeffail/tunny"
)

func main() {
	pool := tunny.NewFunc(3, func(i interface{}) interface{} {
		log.Println(i)
		time.Sleep(time.Second)
		return nil
	})
	defer pool.Close()

	for i := 0; i < 10; i++ {
		go pool.Process(i)
	}
	time.Sleep(time.Second * 4)
}

/*
2024/09/26 23:21:33 0
2024/09/26 23:21:33 1
2024/09/26 23:21:33 3
2024/09/26 23:21:34 2
2024/09/26 23:21:34 6
2024/09/26 23:21:34 4
2024/09/26 23:21:35 7
2024/09/26 23:21:35 5
2024/09/26 23:21:35 8
2024/09/26 23:21:36 9
*/
