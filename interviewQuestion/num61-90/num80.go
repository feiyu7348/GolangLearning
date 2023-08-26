// @author:zfy
// @date:2023/8/26 13:02

package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Printf("i:%d\n", i)
			wg.Done()
		}(wg, i)
	}

	wg.Wait()

	fmt.Println("exit")
}
