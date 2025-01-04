// Package CreationalPattern
// author: zfy  date: 2024/12/6
package CreationalPattern

import "fmt"

// 抽象工厂模式 https://developer.aliyun.com/article/1271217

type Widget interface {
	Display()
}

type Button struct {
}

func (b *Button) Display() {
	fmt.Println("显示按钮")
}

type TextBox struct {
}

func (t *TextBox) Display() {
	fmt.Println("显示文本框")
}

type GUIFactory interface {
	CreateButton() Widget
	CreateTextBox() Widget
}
type WindowsFactory struct {
}

func (w *WindowsFactory) CreateButton() Widget {
	return &Button{}
}

func (w *WindowsFactory) CreateTextBox() Widget {
	return &TextBox{}
}

type MacFactory struct {
}

func (m *MacFactory) CreateButton() Widget {
	return &Button{}
}

func (m *MacFactory) CreateTextBox() Widget {
	return &TextBox{}
}

func main3() {
	windowsFactory := &WindowsFactory{}
	windowsButton := windowsFactory.CreateButton()
	windowsTextBox := windowsFactory.CreateTextBox()
	windowsButton.Display()
	windowsTextBox.Display()

	macFactory := &MacFactory{}
	macButton := macFactory.CreateButton()
	macTextBox := macFactory.CreateTextBox()
	macButton.Display()
	macTextBox.Display()
}
