package main

import (
	"fmt"

	"sync"
)

var mychan = make(chan int, 500)
var done2 = make(chan struct{})
var wg = sync.WaitGroup{}

func main() {

	wg.Add(2)
	go dosth()
	go func() {
		for i := 1; i <= 10000; i++ {
			mychan <- i
		}
		close(done2)
		wg.Done()
	}()
	wg.Wait()
}

func dosth() {
	for {
		fmt.Println(len(mychan))
		select {
		case a := <-mychan:
			//fmt.Println(a)
			a++
		case <-done2:
			goto label
		}

	}
label:
	wg.Done()
}
