// Package StructuralPattern
// author: zfy  date: 2024/12/6
package StructuralPattern

import "fmt"

// 桥接模式 https://developer.aliyun.com/article/1267990

// 抽象部分接口
type Shape interface {
	Draw()
}

// 实现部分接口
type Color interface {
	Fill()
}

// 具体抽象部分：矩形
type Rectangle struct {
	color Color
}

func (r *Rectangle) Draw() {
	fmt.Print("绘制矩形，")
	r.color.Fill()
}

// 具体抽象部分：圆形
type Circle struct {
	color Color
}

func (c *Circle) Draw() {
	fmt.Print("绘制圆形，")
	c.color.Fill()
}

// 具体实现部分：红色
type RedColor struct {
}

func (r *RedColor) Fill() {
	fmt.Println("使用红色填充")
}

// 具体实现部分：蓝色
type BlueColor struct {
}

func (b *BlueColor) Fill() {
	fmt.Println("使用蓝色填充")
}

// 客户端代码
func main() {
	redColor := &RedColor{}
	blueColor := &BlueColor{}

	rectangle := &Rectangle{redColor}
	rectangle.Draw()

	circle := &Circle{blueColor}
	circle.Draw()
}
