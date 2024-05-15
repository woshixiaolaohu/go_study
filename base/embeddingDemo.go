package base

import "fmt"

type notifier1 interface {
	notify()
}

type user1 struct {
	name  string
	email string
}

type admin1 struct {
	user1 //嵌入类型
	lever string
}

func (u *user1) notify() {
	fmt.Printf("Sending email from jelly@gmail.com to %s<%s>\n", u.name, u.email)
}

func (a admin1) notify() {
	fmt.Printf("Admin sending email from jelly@gmail.com to %s<%s>\n", a.name, a.email)
}

func sendNotification1(n notifier1) {
	n.notify()
}

func init() {
	fmt.Printf("\033[1;33;40m%s\033[0m\n", "embeddingDemo")
	a := &admin1{
		user1{
			"user1", "user1@gmail.com",
		}, "high",
	}
	// 外部类型实现了相同的方法时，内部类型的方法不会提升
	// 直接访问内部类型的方法
	a.user1.notify()
	// 内部类型的方法被提升到外部类型
	a.notify()
	sendNotification(a)
}
