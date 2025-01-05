// Package BehavioralPattern
// author: zfy  date: 2024/12/6
package BehavioralPattern

import "fmt"

// 状态模式 https://developer.aliyun.com/article/1268031

// ElevatorContext 环境接口
type ElevatorContext interface {
	SetState(state ElevatorState)
	RequestFloor(floor int)
}

// ElevatorState 抽象状态接口
type ElevatorState interface {
	HandleRequest(context ElevatorContext, floor int)
}

// StoppedState 具体状态：停止状态
type StoppedState struct{}

func (s *StoppedState) HandleRequest(context ElevatorContext, floor int) {
	fmt.Println("电梯已停止，准备前往目标楼层：", floor)
	context.SetState(&MovingState{})
}

// MovingState 具体状态：运行状态
type MovingState struct{}

func (s *MovingState) HandleRequest(context ElevatorContext, floor int) {
	fmt.Println("电梯正在运行，前往目标楼层：", floor)
	context.SetState(&StoppedState{})
}

// ErrorState 具体状态：故障状态
type ErrorState struct{}

func (s *ErrorState) HandleRequest(context ElevatorContext, floor int) {
	fmt.Println("电梯发生故障，请联系维修人员")
}

// Elevator 具体环境
type Elevator struct {
	state ElevatorState
}

func (e *Elevator) SetState(state ElevatorState) {
	e.state = state
}

func (e *Elevator) RequestFloor(floor int) {
	e.state.HandleRequest(e, floor)
}

// 客户端代码
func main() {
	elevator := &Elevator{}
	elevator.SetState(&StoppedState{})

	elevator.RequestFloor(5)
	elevator.RequestFloor(10)

	elevator.SetState(&ErrorState{})
	elevator.RequestFloor(8)
}
