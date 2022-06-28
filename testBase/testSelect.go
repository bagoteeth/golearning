package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1
	c <- 2
	for {
		select {
		case a := <-c:
			fmt.Println(a)
		default:
			fmt.Println("out...")
			return
		}
	}
}
