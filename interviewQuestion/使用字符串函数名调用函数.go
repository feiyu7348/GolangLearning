// @author:zfy
// @date:2024/8/14 11:06

package main

import (
	"fmt"
	"reflect"
)

type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("Eat")
}

func main() {
	a := Animal{}
	reflect.ValueOf(&a).MethodByName("Eat").Call([]reflect.Value{})
}
