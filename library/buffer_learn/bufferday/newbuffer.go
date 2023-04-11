// @author:zfy
// @date:2023/4/11 17:20

package bufferday

import (
	"bytes"
	"fmt"
)

func NewBufferDay() {
	buffer := bytes.NewBuffer(nil)
	fmt.Println(buffer)
	n, err := buffer.WriteString("this is a test for bytes buffer")
	fmt.Println(n, err) // 31  nil
	fmt.Println(buffer)
	fmt.Println(buffer.Len(), buffer.Cap()) // 31 64

	s := make([]byte, 1000)
	n, err = buffer.Read(s)
	fmt.Println(n, err)                     // 31 nil
	fmt.Println(string(s))                  // this is a test for bytes buffer
	fmt.Println(buffer.Len(), buffer.Cap()) // 0 64
}
