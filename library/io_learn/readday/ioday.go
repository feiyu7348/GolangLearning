// @author:zfy
// @date:2023/4/10 14:16

package readday

import (
	"fmt"
	"os"
)

func ReadDay() {
	file, err := os.Open("E:/Github/GolangLearning/library/io_learn/1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	data := make([]byte, 11)
	count, err := file.Read(data)
	if err != nil {
		return
	}
	fmt.Println("字节数据：", data)
	fmt.Println("字符串数据：", string(data)) // Hello world
	fmt.Println("所获取字节的长度：", count)
}
