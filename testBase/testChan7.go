package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	go func() {
		for i := 1; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	time.Sleep(1 * time.Second)
	for i := 1; i < 10; i++ {
		a, ok := <-ch
		fmt.Printf("%d,%t\n", a, ok)
	}
}
