// Package BehavioralPattern
// author: zfy  date: 2024/12/6
package BehavioralPattern

import "fmt"

// 模板方法模式 https://developer.aliyun.com/article/1268036

// CoffeeMaker 抽象类
type CoffeeMaker interface {
	Prepare()
	Brew()
	Serve()
}

// Americano 具体类：美式咖啡
type Americano struct {
}

func (a *Americano) Prepare() {
	fmt.Println("准备美式咖啡的材料")
}

func (a *Americano) Brew() {
	fmt.Println("冲泡美式咖啡")
}

func (a *Americano) Serve() {
	fmt.Println("提供美式咖啡")
}

// Latte 具体类：拿铁咖啡
type Latte struct {
}

func (l *Latte) Prepare() {
	fmt.Println("准备拿铁咖啡的材料")
}

func (l *Latte) Brew() {
	fmt.Println("冲泡拿铁咖啡")
}

func (l *Latte) Serve() {
	fmt.Println("提供拿铁咖啡")
}

// 客户端代码
func main() {
	americano := &Americano{}
	americano.Prepare()
	americano.Brew()
	americano.Serve()

	latte := &Latte{}
	latte.Prepare()
	latte.Brew()
	latte.Serve()
}
