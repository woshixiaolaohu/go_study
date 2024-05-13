package base

import "fmt"

/*
*
方法给用户定义的类型添加行为，实际上也是函数
*/
type user struct {
	name  string
	age   int
	email string
}

// notify 使用值接收者实现一个方法
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// changeEmail 使用指针接收实现一个方法   指针为了修改原值
func (u *user) changeEmail(email string) {
	u.email = email
}
func init() {
	bill := user{
		"bill", 18, "bill@gmail.con",
	}
	bill.notify()
	lisa := &user{
		"lisa", 18, "lisa@gmail.con",
	}
	lisa.notify()
	bill.changeEmail("bill@163.com")
	bill.notify()
}
