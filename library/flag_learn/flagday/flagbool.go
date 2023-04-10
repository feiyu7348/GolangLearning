// author:zfy  date:2023/3/27

package flagday

import (
	"flag"
	"fmt"
)

func FlagBool() {
	boolVar := flag.Bool("testBool", false, "testBool is bool type")
	// 命令行参数进行解析
	flag.Parse()
	fmt.Println(boolVar)
}
