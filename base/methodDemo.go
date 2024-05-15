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
type userFunc struct {
	name      string
	printName func()
}

// Animal 接口类型
type Animal interface {
	GetName() string
	GetAge() int
}

// notify 使用值接收者实现一个方法
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}
func (u user) GetName() string {
	return u.name
}
func (u user) GetAge() int {
	return u.age
}
func f100() {
	name := "testUser"
	fmt.Printf("name: %s\n", name)
}

// changeEmail 使用指针接收实现一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}
func init() {
	fmt.Printf("\033[1;33;40m%s\033[0m\n", "methodDemo")
	testUser := userFunc{
		"testUser",
		f100,
	}
	testUser.printName()
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
	var ani Animal
	// 使用值接收类型会让方法自动拥有指针接收类型方法的能力 指针接收类型则不会
	ani = user{
		name: "wx",
		age:  26,
	}
	ani.GetName()
	ani.GetAge()
}
