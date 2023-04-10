// @author:zfy
// @date:2023/4/10 14:47

package writeday

import (
	"fmt"
	"os"
)

func FeekDay() {
	file, err := os.OpenFile("E:/Github/GolangLearning/library/io_learn/test.txt",
		os.O_CREATE, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.WriteString("G0lang")
	if err != nil {
		return
	}

	_, err = file.Seek(1, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file.Write([]byte{'o'})
	if err != nil {
		fmt.Println(err)
		return
	}
}
