// @author:zfy
// @date:2023/4/10 14:51

package writeday

import (
	"bufio"
	"fmt"
	"os"
)

func Bufioday() {
	file, err := os.OpenFile("E:/Github/GolangLearning/library/io_learn/test.txt",
		os.O_CREATE, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	//fmt.Printf("file: %v\n", file)
	writer := bufio.NewWriter(file)
	//fmt.Printf("writer: %v\n", writer)
	a, err := writer.WriteString("Hello World\n")
	fmt.Printf("a: %v\n", a)
	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := writer.WriteString("Hello Golang\n")
	fmt.Printf("b: %v\n", b)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = writer.WriteString("Hello Gopher\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	writer.Flush()
}
