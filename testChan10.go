package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("end of goroutine1")
				return
			default:
				fmt.Println("goroutine1 is running")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("end of goroutine2")
				return
			default:
				fmt.Println("goroutine2 is running")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()
	for {

	}
}
