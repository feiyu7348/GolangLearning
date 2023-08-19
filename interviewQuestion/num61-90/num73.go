// @author:zfy
// @date:2023/8/19 9:48

package main

func test1() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	return funs
}

func main() {
	funs := test1()
	for _, f := range funs {
		f()
	}
}
