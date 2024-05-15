package base

import "fmt"

func f1(x int, y int) (ret int) {
	return x * y
}

func f2(x int, y int) {
	fmt.Printf("x: %d, y: %d\n", x, y)
}

func f3() {
	fmt.Println("没有参数且没有返回值")
}

func f4() (ret int) {
	return 99
}

func f5(x int, y int) (ret int) {
	ret = x * y
	return
}

func f6() (int, string) {
	return 1, "jelly"
}

func f7(x, y, z int, a, b, c string, e, f bool) (ret int) {
	return x + y + z
}

// 可变长参数必须放在函数参数的最后
func f8(x int, y ...int) {
	fmt.Printf("可变长参数y类型为%T，值为%d\n", y, y)
}

// 函数作为函数参数
func f9(f func() int) {
	ret := f()
	fmt.Printf("函数作为参数打印出来的值为%v，类型为%T\n", ret, ret)
}

// 函数作为函数的返回值
func f10(f func() int) func(int, int) int {
	return f5
}

// 闭包：闭包是一个函数，这个函数包含他外部作用域的变量
func f13(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
func init() {
	fmt.Printf("\033[1;33;40m%s\033[0m\n", "funcDemo")
	deferDemo()
	f8(10, 1, 2, 3, 4, 5)
	f9(f4)
	f11 := f10(f4)
	fmt.Printf("f11的值:%v\n", f11)
	fmt.Printf("f711的类型：%T\n", f11)
	// 匿名函数
	f12 := func(x, y int) {
		fmt.Printf("匿名函数执行结果为%d\n", x+y)
	}
	f12(1, 98)
	// 自执行函数：匿名函数定义之后加()直接执行
	func(x, y int) {
		fmt.Printf("自执行函数执行结果为%d\n", x+y)
	}(1, 99)

	f14 := f13(1)
	ret := f14(2)
	fmt.Printf("ret值为%d\n", ret)

}
func deferDemo() {
	defer fmt.Println("first")
	defer fmt.Println("second")
	fmt.Println("third")
}
