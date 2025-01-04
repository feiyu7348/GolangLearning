// Package StructuralPattern
// author: zfy  date: 2024/12/6
package StructuralPattern

import "fmt"

// 组合模式

// 组件接口
type Employee interface {
	GetName() string
	GetSalary() float64
}

// 叶子对象：普通员工
type Worker struct {
	name   string
	salary float64
}

func (w *Worker) GetName() string {
	return w.name
}

func (w *Worker) GetSalary() float64 {
	return w.salary
}

// 组合对象：经理
type Manager struct {
	name     string
	salary   float64
	substaff []Employee
}

func (m *Manager) GetName() string {
	return m.name
}

func (m *Manager) GetSalary() float64 {
	return m.salary
}

func (m *Manager) AddEmployee(e Employee) {
	m.substaff = append(m.substaff, e)
}

func (m *Manager) RemoveEmployee(e Employee) {
	for i, emp := range m.substaff {
		if emp == e {
			m.substaff = append(m.substaff[:i], m.substaff[i+1:]...)
			break
		}
	}
}

// 客户端代码
func main() {
	worker1 := &Worker{name: "John", salary: 5000}
	worker2 := &Worker{name: "Alice", salary: 4000}

	manager := &Manager{name: "Tom", salary: 10000}
	manager.AddEmployee(worker1)
	manager.AddEmployee(worker2)

	fmt.Println("Manager:", manager.GetName(), "Salary:", manager.GetSalary())

	for _, emp := range manager.substaff {
		fmt.Println("Substaff:", emp.GetName(), "Salary:", emp.GetSalary())
	}
}
