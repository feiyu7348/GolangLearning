// Package BehavioralPattern
// author: zfy  date: 2024/12/6
package BehavioralPattern

import "fmt"

// 访问者模式 https://developer.aliyun.com/article/1268045

type Visitor interface {
	VisitBook(book *Book)
	VisitElectronicProduct(product *ElectronicProduct)
}

type Element interface {
	Accept(visitor Visitor)
}

type ShoppingCartVisitor struct {
	TotalPrice float64
}

func (v *ShoppingCartVisitor) VisitBook(book *Book) {
	v.TotalPrice += book.Price
}

func (v *ShoppingCartVisitor) VisitElectronicProduct(product *ElectronicProduct) {
	v.TotalPrice += product.Price
}

type Book struct {
	Name  string
	Price float64
}

func (b *Book) Accept(visitor Visitor) {
	visitor.VisitBook(b)
}

type ElectronicProduct struct {
	Name  string
	Price float64
}

func (p *ElectronicProduct) Accept(visitor Visitor) {
	visitor.VisitElectronicProduct(p)
}

type ShoppingCart struct {
	Items []Element
}

func (cart *ShoppingCart) AddItem(item Element) {
	cart.Items = append(cart.Items, item)
}

func (cart *ShoppingCart) CalculateTotalPrice(visitor Visitor) float64 {
	for _, item := range cart.Items {
		item.Accept(visitor)
	}
	return visitor.(*ShoppingCartVisitor).TotalPrice
}

func main() {
	cart := &ShoppingCart{}

	book := &Book{Name: "Design Patterns", Price: 49.99}
	electronicProduct := &ElectronicProduct{Name: "Smartphone", Price: 599.99}

	cart.AddItem(book)
	cart.AddItem(electronicProduct)

	visitor := &ShoppingCartVisitor{}
	totalPrice := cart.CalculateTotalPrice(visitor)

	fmt.Printf("Total price: $%.2f\n", totalPrice)
}
