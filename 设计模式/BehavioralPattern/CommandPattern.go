// Package BehavioralPattern
// author: zfy  date: 2024/12/6
package BehavioralPattern

import "fmt"

// 命令模式 https://developer.aliyun.com/article/1268019

// Command 命令接口
type Command interface {
	Execute()
}

// TVOnCommand 具体命令：打开电视
type TVOnCommand struct {
	tv *TV
}

func NewTVOnCommand(tv *TV) *TVOnCommand {
	return &TVOnCommand{
		tv: tv,
	}
}

func (c *TVOnCommand) Execute() {
	c.tv.On()
}

// TVOffCommand 具体命令：关闭电视
type TVOffCommand struct {
	tv *TV
}

func NewTVOffCommand(tv *TV) *TVOffCommand {
	return &TVOffCommand{
		tv: tv,
	}
}

func (c *TVOffCommand) Execute() {
	c.tv.Off()
}

// TV 接收者：电视
type TV struct {
	isOn bool
}

func (t *TV) On() {
	t.isOn = true
	fmt.Println("TV is on")
}

func (t *TV) Off() {
	t.isOn = false
	fmt.Println("TV is off")
}

// RemoteControl 调用者：遥控器
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

// 客户端代码
func main() {
	tv := &TV{}
	tvOnCommand := NewTVOnCommand(tv)
	tvOffCommand := NewTVOffCommand(tv)

	remoteControl := &RemoteControl{}
	remoteControl.SetCommand(tvOnCommand)
	remoteControl.PressButton()

	remoteControl.SetCommand(tvOffCommand)
	remoteControl.PressButton()
}
