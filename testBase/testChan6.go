package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(chOnlySend chan<- int) {
		for i := 1; i <= 5; i++ {
			chOnlySend <- i
		}
	}(ch)
	go func(chOnlyReceive <-chan int) {
		for {
			a := <-chOnlyReceive
			fmt.Println(a)
		}
	}(ch)
	for {

	}
}
