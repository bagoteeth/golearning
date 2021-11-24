package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	for i := 1; ; i++ {
		select {
		case ch <- i:
			fmt.Println("send ", i, " in select, length is ", len(ch))
			if i >= 20 {
				ch = nil
				fmt.Printf("********\n%T\n%+v\n", ch, ch)
			}
		case x := <-ch:
			fmt.Println("receive ", x, " from ch, length is ", len(ch))
		default:
			fmt.Println("do nothing")
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}
