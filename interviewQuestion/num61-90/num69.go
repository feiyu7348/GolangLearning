// @author:zfy
// @date:2023/8/15 21:58

package main

func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

func main() {
	a, b := test(100)
	a()
	b()
}
