// author:zfy  date:2022/6/26

package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine ending...")

		fmt.Println("goroutine running...")

		c <- 666
	}()

	num := <-c
	fmt.Printf("num = %d\n", num)
	fmt.Printf("main goroutine ended")
}
