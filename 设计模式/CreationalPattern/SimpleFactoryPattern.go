// Package CreationalPattern
// author: zfy  date: 2025/1/4
package CreationalPattern

import "fmt"

//简单go工厂模式https://developer.aliyun.com/article/1271214

type Shape interface {
	Draw()
}

type Circle struct {
}

func (c *Circle) Draw() {
	fmt.Println("绘制圆形")
}

type Rectangle struct {
}

func (r *Rectangle) Draw() {
	fmt.Println("绘制矩形")
}

type ShapeFactory struct {
}

func (sf *ShapeFactory) CreateShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return &Circle{}
	case "rectangle":
		return &Rectangle{}
	default:
		return nil
	}
}

func main1() {
	factory := &ShapeFactory{}

	// 创建圆形对象
	circle := factory.CreateShape("circle")
	circle.Draw()

	// 创建矩形对象
	rectangle := factory.CreateShape("rectangle")
	rectangle.Draw()
}
