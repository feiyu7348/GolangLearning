// Package BehavioralPattern
// author: zfy  date: 2024/12/6
package BehavioralPattern

import "fmt"

// 中介者模式 https://developer.aliyun.com/article/1268025

// Mediator 中介者接口
type Mediator interface {
	SendMessage(message string, user User)
}

// ChatRoom 具体中介者：聊天室
type ChatRoom struct{}

func (cr *ChatRoom) SendMessage(message string, user User) {
	fmt.Printf("%s 发送消息：%s\n", user.GetName(), message)
}

// User 同事对象接口
type User interface {
	GetName() string
	SendMessage(message string)
}

// ChatUser 具体同事对象：用户
type ChatUser struct {
	name     string
	mediator Mediator
}

func NewChatUser(name string, mediator Mediator) *ChatUser {
	return &ChatUser{
		name:     name,
		mediator: mediator,
	}
}

func (cu *ChatUser) GetName() string {
	return cu.name
}

func (cu *ChatUser) SendMessage(message string) {
	cu.mediator.SendMessage(message, cu)
}

// 客户端代码
func main() {
	chatRoom := &ChatRoom{}

	user1 := NewChatUser("User1", chatRoom)
	user2 := NewChatUser("User2", chatRoom)
	user3 := NewChatUser("User3", chatRoom)

	user1.SendMessage("Hello, User2!")
	user2.SendMessage("Hi, User1!")
	user3.SendMessage("Nice to meet you all!")
}
