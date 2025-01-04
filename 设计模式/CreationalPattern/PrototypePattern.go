// Package CreationalPattern
// author: zfy  date: 2024/12/6
package CreationalPattern

import "fmt"

// 原型模式 https://developer.aliyun.com/article/1267726

type Shape1 interface {
	Clone() Shape1
	Draw()
}

type Rectangle1 struct {
}

func (r *Rectangle1) Clone() Shape1 {
	return &Rectangle1{}
}

func (r *Rectangle1) Draw() {
	fmt.Println("Drawing a rectangle")
}

type Circle1 struct {
}

func (c *Circle1) Clone() Shape1 {
	return &Circle1{}
}

func (c *Circle1) Draw() {
	fmt.Println("Drawing a circle")
}

func main4() {
	rectangle := &Rectangle1{}
	clonedRectangle := rectangle.Clone()
	rectangle.Draw()
	clonedRectangle.Draw()

	circle := &Circle1{}
	clonedCircle := circle.Clone()
	circle.Draw()
	clonedCircle.Draw()
}
