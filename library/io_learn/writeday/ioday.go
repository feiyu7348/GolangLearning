// @author:zfy
// @date:2023/4/10 14:42

package writeday

import (
	"fmt"
	"os"
)

func WriteDay() {
	file, err := os.OpenFile("E:/Github/GolangLearning/library/io_learn/1.txt",
		os.O_CREATE, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	count, err := file.Write([]byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\n'})
	if err != nil {
		return
	}
	fmt.Printf("写入了 %d 字节\n", count)

	count, err = file.WriteString("Hello Golang")
	if err != nil {
		return
	}
	fmt.Printf("写入了长度为 %d 的字符串\n", count)

	count, err = file.WriteAt([]byte{'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x'}, 2)
	if err != nil {
		return
	}
	fmt.Printf("写入了 %d 字节\n", count)
}
