// @author:zfy
// @date:2023/8/10 22:17

package main

func main() {
	defer func() {
		recover()
	}()
	panic(1)
}

// recover() 必须在 defer() 函数中直接调用才有效
