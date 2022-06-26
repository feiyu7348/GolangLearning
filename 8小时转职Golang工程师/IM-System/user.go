// author:zfy  date:2022/6/26

package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}
