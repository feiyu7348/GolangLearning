// @author:zfy
// @date:2023/7/29 11:50

package main

type X struct{}

func (x *X) test() {
	println(x)
}

func main() {
	var a *X
	a.test() // 相当于 test(nil)
	var x = X{}
	x.test()
}

//0x0
//0xc000055f70
