// author:zfy  date:2023/3/28

package regexday

import (
	"fmt"
	"log"
	"regexp"
)

func CompileDay() {
	words := [...]string{"1", "2", "3", "4", "5", "6"}
	// 编译的正则表达式
	cpu := "0-5"
	//a := "[cpu]"
	re, err := regexp.Compile(cpu)
	if err != nil {
		log.Fatal(err)
	}

	for _, word := range words {
		// 在返回的正则表达式对象上调用 MatchString 函数
		found := re.MatchString(word)
		if found {
			fmt.Printf("%s matches\n", word)
		} else {
			fmt.Printf("%s does not match\n", word)
		}
	}
}
