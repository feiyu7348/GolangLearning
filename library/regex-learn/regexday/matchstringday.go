// author:zfy  date:2023/3/28

package regexday

import (
	"fmt"
	"log"
	"regexp"
)

func MatchString() {
	words := [...]string{"Seven", "even", "Maven", "Amen", "eleven"}
	for _, word := range words {
		found, err := regexp.MatchString(".even", word)
		if err != nil {
			log.Fatal(err)
		}

		if found {
			fmt.Printf("%s matches\n", word)
		} else {
			fmt.Printf("%s dose not match\n", word)
		}
	}
}

/* 输出
Seven matches
even dose not match
Maven dose not match
Amen dose not match
eleven matches
*/
