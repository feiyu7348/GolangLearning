// author:zfy  date:2022/6/26

package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go newTask()

	i := 0
	for {
		i++
		fmt.Printf("main Goroutine : i = %d", i)
		time.Sleep(1 * time.Second)
	}
}
