package base

import "fmt"

// 接口
// notifier 通知行为类接口
type notifier interface {
	notify()
}

type person struct {
	name  string
	email string
}

func (p *person) notify() {
	fmt.Printf("Sending user email to%s<%s>\n", p.name, p.email)
}

func (p *person) log() {
	fmt.Println("p.log func")
}

func sendNotification(n notifier) {
	n.notify()
}

func init() {
	u := &person{
		"bilibili", "vill@gmail.com",
	}
	u.log()
	sendNotification(u)
}
