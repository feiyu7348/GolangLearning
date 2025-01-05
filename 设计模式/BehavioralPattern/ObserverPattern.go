// Package BehavioralPattern
// author: zfy  date: 2024/12/6
package BehavioralPattern

import "fmt"

// 观察者模式 https://developer.aliyun.com/article/1267681

// Subject 主题接口
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify()
}

// Observer 观察者接口
type Observer interface {
	Update(subject Subject)
}

// NewsSubject 具体主题
type NewsSubject struct {
	observers []Observer
	news      string
}

func (n *NewsSubject) Register(observer Observer) {
	n.observers = append(n.observers, observer)
}

func (n *NewsSubject) Unregister(observer Observer) {
	for i, o := range n.observers {
		if o == observer {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			break
		}
	}
}

func (n *NewsSubject) Notify() {
	for _, observer := range n.observers {
		observer.Update(n)
	}
}

func (n *NewsSubject) SetNews(news string) {
	n.news = news
	n.Notify()
}

// Subscriber 具体观察者
type Subscriber struct {
	name string
}

func (s *Subscriber) Update(subject Subject) {
	newsSubject := subject.(*NewsSubject)
	fmt.Printf("[%s] 收到新闻通知：%s\n", s.name, newsSubject.news)
}

// 客户端代码
func main() {
	subject := &NewsSubject{}

	subscriber1 := &Subscriber{name: "订阅者1"}
	subject.Register(subscriber1)

	subscriber2 := &Subscriber{name: "订阅者2"}
	subject.Register(subscriber2)

	subject.SetNews("新闻1发布了")
	subject.SetNews("新闻2发布了")

	subject.Unregister(subscriber1)

	subject.SetNews("新闻3发布了")
}
