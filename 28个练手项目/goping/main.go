// @author:zfy
// @date:2023/4/7 23:37

package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	timeOut int64
	size    int
	count   int
	typ     uint8 = 8
	code    uint8 = 8
)

type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	ID          uint16
	SequenceNum uint16
}

func getCommandArgs() {
	flag.Int64Var(&timeOut, "w", 1000, "请求超时时长，单位毫秒")
	flag.IntVar(&size, "l", 32, "请求发送缓冲区大小，单位字节")
	flag.IntVar(&count, "n", 4, "发生请求数")
	flag.Parse()
}

func checkSum(data []byte) uint16 {
	length := len(data)
	index := 0
	var sum uint32
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		length -= 2
		index += 2
	}

	if length != 0 {
		sum += uint32(data[index])
	}

	hi16 := sum >> 16
	for hi16 != 0 {
		sum = hi16 + uint32(uint16(sum))
		hi16 = sum >> 16
	}

	return uint16(^sum)
}

func main() {
	getCommandArgs()
	desIp := os.Args[len(os.Args)-1]
	conn, err := net.DialTimeout("ip:icmp", desIp, time.Duration(timeOut)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	icmp := &ICMP{
		Type:        typ,
		Code:        code,
		CheckSum:    0,
		ID:          1,
		SequenceNum: 1,
	}
	data := make([]byte, size)
	fmt.Println(data)
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, icmp)
	buf.Write(data)
	data = buf.Bytes()
	checkSum := checkSum(data)
	data[2] = byte(checkSum >> 8)
	data[3] = byte(checkSum)
	conn.SetDeadline(time.Now().Add(time.Duration(timeOut) * time.Millisecond))
	n, err := conn.Write(data)
	if err != nil {
		log.Println()
		return
	}
	buffer := make([]byte, 65535)
	n, err = conn.Read(buffer)
	if err != nil {
		log.Println()
		return
	}
	fmt.Println(n)
}
