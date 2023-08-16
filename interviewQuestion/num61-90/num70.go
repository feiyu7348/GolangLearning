// @author:zfy
// @date:2023/8/16 20:13

package main

func main() {

	println(DeferTest1(1))
	println(DeferTest2(1))
}

func DeferTest1(i int) (r int) {
	r = i
	defer func() {
		r += 3
	}()
	return r
}

func DeferTest2(i int) (r int) {
	defer func() {
		r += i
	}()
	return 2
}
