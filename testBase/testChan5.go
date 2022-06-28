package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 3
	}()
	go func() {
		ch <- 2
		ch <- 4
	}()
	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)

	}()
	time.Sleep(1 * time.Second)
}
