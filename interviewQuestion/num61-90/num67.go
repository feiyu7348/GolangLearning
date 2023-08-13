// @author:zfy
// @date:2023/8/13 9:31

package main

import "fmt"

func main() {
	ts := [2]T{}
	for i := range ts[:] {
		switch i {
		case 0:
			ts[1].n = 9
		case 1:
			fmt.Print(ts[i].n, " ")
		}
	}
	fmt.Print(ts)

	//ts := [2]T{}
	//for i := range ts[:] {
	//	switch t := &ts[i]; i {
	//	case 0:
	//		t.n = 3;
	//		ts[1].n = 9
	//	case 1:
	//		fmt.Print(t.n, " ")
	//	}
	//}
	//fmt.Print(ts)
}
