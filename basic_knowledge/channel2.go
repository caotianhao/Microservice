package main

import "fmt"

func main() {
	numChan := make(chan int)
	go func() {
		for i := 0; i < 33; i++ {
			numChan <- i * i
			fmt.Println("write", i*i)
		}
		// for-range 不知道什么时候结束，所以写入之后要关闭
		close(numChan)
	}()

	// for-range 遍历 channel 时只有一个参数
	for i := range numChan {
		fmt.Println("-------------read", i)
	}
}
