package main

import (
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func main() {
	wg2.Add(2)
	go func() {
		defer wg2.Done()
		time.Sleep(5 * time.Second)
		wg2.Add(1)
	}()
	go func() {
		defer wg2.Done()
		time.Sleep(3 * time.Second)
	}()
	wg2.Wait()
}
