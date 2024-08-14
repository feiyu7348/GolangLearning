// @author:zfy
// @date:2024/8/14 11:21

package main

import (
	"fmt"
	"sync"
)

// 两个协程交替打印10个字母和数字
// 思路：采用channel来协调goroutine之间顺序。
//主线程一般要waitGroup等待协程退出，这里简化了一下直接sleep。

func main() {
	letter, number := make(chan bool), make(chan bool) // 两个 chan 用于交替执行两个子协程
	wait := sync.WaitGroup{}                           // 设置计数器，用于控制主协程堵塞等待子协程执行
	go func() {
		i := 1
		for {
			select {
			case <-number: // 堵塞，等待number有值后向下执行
				fmt.Print(i) // 打印两个连续的数字
				i++
				fmt.Print(i)
				i++
				letter <- true // 赋值后，打印字母协程收到后继续执行
			}
		}
	}()
	wait.Add(1) // 计数器加1
	go func() {
		i := 'A'
		for {
			select {
			case <-letter: // 堵塞，等待letter有值向下执行
				if i >= 'Z' { // 子协程结束
					wait.Done() // 计数器设置为0，退出主线程
					return
				}
				fmt.Print(string(i)) // 打印两个连续的字母
				i++
				fmt.Print(string(i))
				i++
				number <- true // 赋值后，打印数字协程收到后继续执行
			}
		}
	}()
	fmt.Println("程序先执行此句输出。")
	number <- true // 赋值后，打印数字协程收到后继续执行
	wait.Wait()    // 堵塞主协程，直到计数器为0
	fmt.Println("\n程序最后执行此句输出。")
}
