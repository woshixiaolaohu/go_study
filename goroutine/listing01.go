package goroutine

import (
	"fmt"
	"runtime"
	"sync"
)

// wg用来等待程序完成
var wg sync.WaitGroup

// 创建goroutine及调度器行为
func init1() {
	fmt.Printf("\033[1;33;40m%s\033[0m\n", "goroutine-listing01")

	// 给每个可用核心分配一个逻辑处理器
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("NumCpu: %d\n", runtime.NumCPU())
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// 计数加2，标示要等待2个goroutine
	wg.Add(4)
	fmt.Println("Start Goroutines")

	go printPrime("First")
	go printPrime("Second")
	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		// 在函数退出时调用Done来通知init函数工作已经完成
		defer wg.Done()
		// 显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		// 函数退时调用Done通知main函数工作已经完成
		defer wg.Done()

		// 显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	// 等待goroutine结束
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

// printPrime 显示5000以内的素数
func printPrime(prefix string) {
	// 函数退出时调用Done告知main函数工作已经完成
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
