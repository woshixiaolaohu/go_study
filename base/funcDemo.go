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

func init() {
	deferDemo()
}
func deferDemo() {
	defer fmt.Println("first")
	defer fmt.Println("second")
	fmt.Println("third")
}
