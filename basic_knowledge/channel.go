package main

import "fmt"

func main() {
	// 无缓冲通道
	//numChan := make(chan int)
	// 有缓冲通道
	numChan := make(chan int, 5)

	// read
	go func() {
		for i := 0; i < 50; i++ {
			r := <-numChan
			fmt.Println("-------------read", r)
		}
	}()

	// write1
	go func() {
		for i := 0; i < 20; i++ {
			numChan <- i
			fmt.Println("write-------", i)
		}
	}()

	// write2
	for i := 20; i < 50; i++ {
		numChan <- i
		fmt.Println("write", i)
	}
}
