// @author:zfy
// @date:2023/4/10 14:55

package writeday

import (
	"os"
)

func AppendDay() {
	// 打开指定的文件
	// 已不存在创建、追加、读写的方式打开
	f, err := os.OpenFile("E:/Github/GolangLearning/library/io_learn/test.txt",
		os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	// 写入追加数据
	f.WriteString("追加的内容\n")
}
