// @author:zfy
// @date:2023/4/10 22:02

package bufferday

import (
	"bytes"
	"fmt"
)

func BufferInit() {
	var buffer bytes.Buffer
	n, err := buffer.WriteString("this is a test for bytes buffer")
	fmt.Println(n, err)                     // 31  nil
	fmt.Println(buffer.Len(), buffer.Cap()) // 31 64

	s := make([]byte, 1000)
	n, err = buffer.Read(s)
	fmt.Println(n, err)                     // 31 nil
	fmt.Println(string(s))                  // this is a test for bytes buffer
	fmt.Println(buffer.Len(), buffer.Cap()) // 0 64
}
