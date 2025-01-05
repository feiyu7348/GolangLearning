// Package BehavioralPattern
// author: zfy  date: 2024/12/6
package BehavioralPattern

import "fmt"

// 备忘录模式 https://developer.aliyun.com/article/1268026

// Memento 备忘录对象
type Memento struct {
	state string
}

func (m *Memento) GetState() string {
	return m.state
}

// Originator 发起人对象
type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{
		state: o.state}
}

func (o *Originator) RestoreMemento(memento *Memento) {
	o.state = memento.GetState()
}

// Caretaker 管理者对象
type Caretaker struct {
	memento *Memento
}

func (c *Caretaker) SaveMemento(memento *Memento) {
	c.memento = memento
}

func (c *Caretaker) GetMemento() *Memento {
	return c.memento
}

// 客户端代码
func main() {
	originator := &Originator{}
	caretaker := &Caretaker{}

	originator.SetState("State 1")
	fmt.Println("当前状态:", originator.state)

	caretaker.SaveMemento(originator.CreateMemento())

	originator.SetState("State 2")
	fmt.Println("当前状态:", originator.state)

	originator.RestoreMemento(caretaker.GetMemento())
	fmt.Println("恢复后的状态:", originator.state)
}
