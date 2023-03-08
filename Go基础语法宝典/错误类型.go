// author:zfy  date:2022/6/26

package main

import (
	"errors"
	"fmt"
)

func main() {
	// Go内置有一个error类型，专门用来处理错误信息，Go的package里面还专门有一个包errors来处理错误
	err := errors.New("emit macho dwarf: elf header corrupted")
	fmt.Print(err)
}
