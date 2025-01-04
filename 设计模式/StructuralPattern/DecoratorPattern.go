// Package StructuralPattern
// author: zfy  date: 2024/12/6
package StructuralPattern

import "fmt"

// 装饰器模式 https://developer.aliyun.com/article/1267998

// 抽象组件：咖啡接口
type Coffee interface {
	GetDescription() string
	GetCost() float64
}

// 具体组件：原味咖啡
type PlainCoffee struct {
}

func (c *PlainCoffee) GetDescription() string {
	return "Plain Coffee"
}

func (c *PlainCoffee) GetCost() float64 {
	return 1.0
}

// 装饰器：配料装饰器
type IngredientDecorator struct {
	coffee    Coffee
	extraCost float64
	extraDesc string
}

func (d *IngredientDecorator) GetDescription() string {
	return d.coffee.GetDescription() + ", " + d.extraDesc
}

func (d *IngredientDecorator) GetCost() float64 {
	return d.coffee.GetCost() + d.extraCost
}

// 具体装饰器：牛奶装饰器
type MilkDecorator struct {
	IngredientDecorator
}

func NewMilkDecorator(coffee Coffee) *MilkDecorator {
	return &MilkDecorator{
		IngredientDecorator{
			coffee:    coffee,
			extraCost: 0.5,
			extraDesc: "Milk",
		},
	}
}

// 具体装饰器：糖浆装饰器
type SyrupDecorator struct {
	IngredientDecorator
}

func NewSyrupDecorator(coffee Coffee) *SyrupDecorator {
	return &SyrupDecorator{
		IngredientDecorator{
			coffee:    coffee,
			extraCost: 0.3,
			extraDesc: "Syrup",
		},
	}
}

// 客户端代码
func main() {
	coffee := &PlainCoffee{}
	fmt.Println("Coffee:", coffee.GetDescription(), "Cost:", coffee.GetCost())

	coffeeWithMilk := NewMilkDecorator(coffee)
	fmt.Println("Coffee with Milk:", coffeeWithMilk.GetDescription(), "Cost:", coffeeWithMilk.GetCost())

	coffeeWithMilkAndSyrup := NewSyrupDecorator(coffeeWithMilk)
	fmt.Println("Coffee with Milk and Syrup:", coffeeWithMilkAndSyrup.GetDescription(), "Cost:", coffeeWithMilkAndSyrup.GetCost())
}
