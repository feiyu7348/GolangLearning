// author:zfy  date:2023/3/28

package regexday

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ZfyExec() {
	args := []string{"eth10", "0", "1-5", "6"}
	cpuList := make([]string, 0)
	re, err := regexp.Compile("\\d-\\d")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range args[1:] {
		found := re.MatchString(v)
		if found {
			fmt.Printf("v: %s\n", v)
			l := strings.Split(v, "-")
			fmt.Printf("l: %s\n", l)
			startInt, _ := strconv.Atoi(l[0])
			endInt, _ := strconv.Atoi(l[1])
			for i := startInt; i <= endInt; i++ {
				cpu := strconv.Itoa(i)
				cpuList = append(cpuList, cpu)
			}
		} else {
			cpuList = append(cpuList, v)
		}
	}

	fmt.Println(cpuList)
}
