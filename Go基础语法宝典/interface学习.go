// author:zfy  date:2022/6/29

package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float64
}

type Employee struct {
	Human
	company string
	money   float64
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on &s\n", h.name, h.phone)
}

func (h *Human) Sing(lyrics string) {
	fmt.Println("la la, la la la, la la la la...", lyrics)
}

func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s, call me on &s\n", e.name, e.company, e.phone)
}

func (s *Student) BorrowMoney(amount float64) {
	s.loan += amount
}

func (e *Employee) SpendSalary(amount float64) {
	e.money -= amount
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float64)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float64)
}
