package main

import (
	"fmt"
	"go_study/base"
	_ "go_study/goroutine"
	//_ "go_study/database"
	"log"
	"os"
)

func init() {
	// 输出标准日志
	log.SetOutput(os.Stdout)
}
func main() {
	fmt.Printf("\033[1;33;40m%s\033[0m\n", "main")
	counter := base.AlertCounter(10)
	counter1 := base.New(100)
	counter2 := base.Counter{Name: "counterName", Value: 999}
	fmt.Printf("Hello counter: %d, New counter1: %d, counter2: %v\n", counter, counter1, counter2.Name)
}
