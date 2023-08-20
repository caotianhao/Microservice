package main

import (
	"fmt"
	"runtime"
	"time"
)

// runtime.Goexit()
// 提前退出 go 程
// 他俩的关系就相当于 for 和 break
func main() {
	go func() {
		func() {
			fmt.Println("1. 子 go 程内部的函数")
			//return // 1234 句全部打印
			//os.Exit(-1) // 24 一定没有，若主 go 程抢到资源则有 3
			runtime.Goexit() // 仅没有 2，符合要求
		}()
		fmt.Println("2. 子 go 程函数")
	}()
	fmt.Println("3. 主 go 程函数")
	time.Sleep(1 * time.Second)
	fmt.Println("4. OVER")
}
