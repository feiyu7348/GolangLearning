// @author:zfy
// @date:2023/4/11 11:32

package netday

import (
	"fmt"
	"net"
)

type parseIPTest struct {
	In  string
	Out net.IP
}

func ParseIP() {
	a := &parseIPTest{
		In: "127.0.1.2",
	}

	a.Out = net.ParseIP(a.In)
	fmt.Printf("%T, %T, %T\n%+v", a, a.In, a.Out, a)
}
