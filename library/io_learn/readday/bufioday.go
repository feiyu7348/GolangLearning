// @author:zfy
// @date:2023/4/10 14:34

package readday

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func BufioDay() {
	file, err := os.OpenFile("E:/Github/GolangLearning/library/io_learn/1.txt",
		os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		if lineData, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				// 因为是以换行符为分隔符，如果最后一行没有换行符，那么返回 io.EOF 错误时，
				// 也是可能读到数据的，因此判断一下是否读到了数据
				if len(lineData) > 0 {
					fmt.Println(lineData)
				}
				break
			}
		} else {
			fmt.Println(strings.TrimRight(lineData, "\n"))
		}
	}
}
