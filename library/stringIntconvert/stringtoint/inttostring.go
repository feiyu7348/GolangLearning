// author:zfy  date:2023/3/29

package stringtoint

import (
	"fmt"
	"strconv"
)

func IntToString() {
	var a int = 1
	var b int64 = 1
	aString := strconv.Itoa(a)

	bString := strconv.FormatInt(b, 64)

	fmt.Println(aString)
	fmt.Println(bString)
}
