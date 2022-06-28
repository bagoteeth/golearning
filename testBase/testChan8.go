package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 9; i++ {
			ch <- i
		}
		fmt.Println("send finished")
	}()
	go func() {
		for i := range ch {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	time.Sleep(3 * time.Second)
}
