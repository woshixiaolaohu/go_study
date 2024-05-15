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

type admin struct {
	name  string
	email string
}

func (p *person) notify() {
	fmt.Printf("Sending user email to%s<%s>\n", p.name, p.email)
}

func (p *person) log() {
	fmt.Println("p.log func")
}

func (a admin) notify() {
	fmt.Printf("Receive Email from %s<%s>\n", a.name, a.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func init() {
	fmt.Printf("\033[1;33;40m%s\033[0m\n", "interfafeDemo")
	u := &person{
		"bilibili", "vill@gmail.com",
	}
	a := &admin{
		"admin", "admin@gmail.com",
	}
	u.log()
	sendNotification(a)
	sendNotification(u)
}
