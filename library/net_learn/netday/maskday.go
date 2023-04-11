// @author:zfy
// @date:2023/4/11 11:50

package netday

import (
	"fmt"
	"net"
)

func MaskDay() {
	var ip = net.IP{127, 168, 124, 1}
	fmt.Println(ip.DefaultMask()) //ff000000,即255.0.0.0
	fmt.Printf("%q\n", ip)        //"127.168.124.1"
	//将子网掩码设为255.255.0.0后，返回的ip将会符合对应的子网掩码，所以返回"127.168.0.0"
	//如果设置的是255.0.0.0，则返回"127.0.0.0"
	ip = ip.Mask(net.IPv4Mask(255, 255, 0, 0))
	fmt.Printf("%q\n", ip) //"127.168.0.0"
}
