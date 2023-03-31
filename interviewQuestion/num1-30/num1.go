// @author:zfy
// @date:2023/3/31 17:45
// http://mian.topgoer.com/%E7%AC%AC%E4%B8%80%E5%A4%A9/

package main

import "fmt"

func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

func main() {
	deferCall()
}

//打印后
//打印中
//打印前
//panic: 触发异常
//参考解析：defer 的执行顺序是后进先出。当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，最后才会执行panic
