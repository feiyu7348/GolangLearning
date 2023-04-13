// @author:zfy
// @date:2023/4/13 10:14

package main

import "fmt"

func hello3(i int) {
	fmt.Println(i)
}

//func main() {
//	i := 5
//	defer hello3(i)
//	i = i + 10
//}

//5

type People1 struct{}

func (p *People1) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People1) ShowB() {
	fmt.Println("showB")
}

type Teacher1 struct {
	People1
}

func (t *Teacher1) ShowB() {
	fmt.Println("teacher showB")
}
func main() {
	t := Teacher1{}
	t.ShowA()
}
