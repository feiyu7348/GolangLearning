// @author:zfy
// @date:2023/8/24 21:51

package main

func alwaysFalse() bool {
	return false
}

func main() {
	switch alwaysFalse(); {
	case true:
		println(true)
	case false:
		println(false)
	}
}
