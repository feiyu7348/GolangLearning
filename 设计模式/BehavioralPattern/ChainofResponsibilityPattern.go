// Package BehavioralPattern
// author: zfy  date: 2024/12/6
package BehavioralPattern

import "fmt"

// 责任链模式 https://developer.aliyun.com/article/1268014

// 抽象处理者：审批者接口
type Approver interface {
	SetNext(approver Approver)
	ProcessRequest(request Request)
}

// 具体处理者：部门经理
type DepartmentManager struct {
	next Approver
}

func (dm *DepartmentManager) SetNext(approver Approver) {
	dm.next = approver
}

func (dm *DepartmentManager) ProcessRequest(request Request) {
	if request.Type == "Leave" && request.Amount <= 3 {
		fmt.Println("Department Manager approved the request.")
	} else if dm.next != nil {
		dm.next.ProcessRequest(request)
	}
}

// 具体处理者：人事部门
type HRDepartment struct {
	next Approver
}

func (hr *HRDepartment) SetNext(approver Approver) {
	hr.next = approver
}

func (hr *HRDepartment) ProcessRequest(request Request) {
	if request.Type == "Leave" && request.Amount <= 7 {
		fmt.Println("HR Department approved the request.")
	} else if hr.next != nil {
		hr.next.ProcessRequest(request)
	}
}

// 具体处理者：总经理
type GeneralManager struct {
	next Approver
}

func (gm *GeneralManager) SetNext(approver Approver) {
	gm.next = approver
}

func (gm *GeneralManager) ProcessRequest(request Request) {
	if request.Type == "Leave" && request.Amount <= 10 {
		fmt.Println("General Manager approved the request.")
	} else {
		fmt.Println("Request denied.")
	}
}

// 请求结构体
type Request struct {
	Type   string
	Amount int
}

// 客户端代码
func main() {
	departmentManager := &DepartmentManager{}
	hrDepartment := &HRDepartment{}
	generalManager := &GeneralManager{}

	departmentManager.SetNext(hrDepartment)
	hrDepartment.SetNext(generalManager)

	request := Request{
		Type:   "Leave",
		Amount: 5,
	}

	departmentManager.ProcessRequest(request)
}
