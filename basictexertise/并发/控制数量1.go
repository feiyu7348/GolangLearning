// Package main
// author: zfy  date: 2024/9/26
package main

import (
	"log"
	"sync"
	"time"
)

type Glimit struct {
	n int
	c chan struct{}
}

// New initialization Glimit struct
func New(n int) *Glimit {
	return &Glimit{
		n: n,
		c: make(chan struct{}, n),
	}
}

// Run f in a new goroutine but with limit.
func (g *Glimit) Run(f func()) {
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
	}()
}

var wg = sync.WaitGroup{}

func main() {
	number := 10
	g := New(2)
	for i := 0; i < number; i++ {
		wg.Add(1)
		value := i
		goFunc := func() {
			// 做一些业务逻辑处理
			log.Printf("go func: %d\n", value)
			time.Sleep(time.Second)
			wg.Done()
		}
		g.Run(goFunc)
	}
	wg.Wait()
}

/*
2024/09/26 23:16:20 go func: 1
2024/09/26 23:16:20 go func: 0
2024/09/26 23:16:21 go func: 2
2024/09/26 23:16:21 go func: 3
2024/09/26 23:16:22 go func: 4
2024/09/26 23:16:22 go func: 5
2024/09/26 23:16:23 go func: 6
2024/09/26 23:16:23 go func: 7
2024/09/26 23:16:24 go func: 8
2024/09/26 23:16:24 go func: 9
*/
