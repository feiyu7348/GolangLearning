// @author:zfy
// @date:2024/8/14 11:15

package main

// Goroutine）有三个函数，分别打印"cat", "fish","dog"要求每一个函数都用一个goroutine，按照顺序打印100次。
// 此题目考察channel，用三个无缓冲channel，如果一个channel收到信号则通知下一个。

import (
	"fmt"
	"time"
)

var dog = make(chan struct{})
var cat = make(chan struct{})
var fish = make(chan struct{})

func Dog() {
	<-fish
	fmt.Println("dog")
	dog <- struct{}{}
}

func Cat() {
	<-dog
	fmt.Println("cat")
	cat <- struct{}{}
}

func Fish() {
	<-cat
	fmt.Println("fish")
	fish <- struct{}{}
}

func main() {
	for i := 0; i < 6; i++ {
		go Dog()
		go Cat()
		go Fish()
	}
	dog <- struct{}{}
	time.Sleep(10 * time.Second)
}
