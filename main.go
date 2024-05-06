package main

import (
	_ "go_study/base"
	"log"
	"os"
)

func init() {
	// 输出标准日志
	log.SetOutput(os.Stdout)
}
func main() {
	//fmt.Println("hello,main.go!")
}
