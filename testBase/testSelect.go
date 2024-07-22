package main

import "fmt"

func main() {
	c := make(chan int)
	select {
	case a, ok := <-c:
		fmt.Printf("receive %d, %b", a, ok)
	}

	//a, ok := <- c
	//fmt.Printf("receive %d, %b", a, ok)
}
