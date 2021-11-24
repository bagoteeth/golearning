package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		ch <- 99 //阻塞，直到5秒后另一个goroutine接收
	}()
	go func() {
		fmt.Println(<-ch)
		fmt.Println("should wait for 5 second")

	}()
	for i := 0; i < 10; i++ { //计时器，每秒打印
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
