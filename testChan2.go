package main

import "fmt"

var mychan = make(chan int, 500)

func main() {
	go dosth()
	go func() {
		for i := 1; i <= 1000; i++ {
			mychan <- i
		}
	}()
}

func dosth() {
	for {
		fmt.Println(len(mychan))
		select {
		case a := <-mychan:
			fmt.Println(a)
		}
	}
}
