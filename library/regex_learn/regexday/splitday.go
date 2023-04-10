// author:zfy  date:2023/3/28

package regexday

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func SplitDay() {
	var data = `22, 1, 3, 4, 5, 17, 4, 3, 21, 4, 5, 1, 48, 9, 42`
	sum := 0
	// \s匹配空格字符
	re := regexp.MustCompile(",\\s*")
	log.Printf("re: %s", re)
	vals := re.Split(data, -1)
	log.Printf("vals: %v", vals)
	for _, val := range vals {
		n, err := strconv.Atoi(val)
		sum += n
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(sum)
}
