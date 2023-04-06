// author:zfy  date:2023/3/29

package stringtoint

import (
	"fmt"
	"log"
	"strconv"
)

func StringToInt() {
	var s = "1"
	sInt, _ := strconv.Atoi(s)

	sInt64, err := strconv.ParseInt(s, 10, 64)
	sUint32 := uint32(sInt64)
	if err != nil {
		log.Println(err)
	}

	sFloat64, _ := strconv.ParseFloat(s, 64)

	fmt.Println(sInt)
	fmt.Println(sInt64)
	fmt.Println(sUint32)
	fmt.Println(sFloat64)
}
