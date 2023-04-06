// @author:zfy
// @date:2023/4/3 16:04

package main

import (
	"bytes"
	"errors"
	"fmt"
)

type fileError struct {
	s string
}

func main() {
	errs := make([]error, 0)
	//errs = append(errs, errors.New("123"))
	//errs = append(errs, errors.New("456"))

	//var errsString string
	//for _, v := range errs {
	//	errsString += v.Error() + "\n"
	//}
	//
	//if errsString != "" {
	//	fmt.Println(errors.New(errsString))
	//} else {
	//	fmt.Println(nil)
	//}

	buf := new(bytes.Buffer)
	fmt.Printf("buf: %s", buf)
	for _, v := range errs {
		buf.WriteString(v.Error() + "\n")
	}

	fmt.Println(errors.New(buf.String()))
	if len(buf.String()) == 0 {
		fmt.Println(nil)
	} else {
		fmt.Println("not nil")
	}
}
