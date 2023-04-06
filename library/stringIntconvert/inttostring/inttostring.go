// author:zfy  date:2023/3/29

package inttostring

import (
	"fmt"
	"strconv"
)

func IntToString() {
	var a int = 1
	var b int64 = 136
	aString := strconv.Itoa(a)

	bString := strconv.FormatInt(b, 10)

	fmt.Println(aString)
	fmt.Println(bString)
}
