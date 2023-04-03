// @author:zfy
// @date:2023/4/3 16:04

package main

import (
	"errors"
	"fmt"
)

type fileError struct {
	s string
}

func main() {
	errs := make([]error, 0)
	errs = append(errs, errors.New("123"))
	errs = append(errs, errors.New("456"))

	var errsString string
	for _, v := range errs {
		errsString += v.Error() + "\n"
	}

	fmt.Println(errsString)
}
