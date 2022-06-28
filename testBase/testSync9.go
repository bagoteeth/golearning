package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make([]int, 20000)
	a[0] = 114514
	//mu:=sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		for i := 1; i <= 10000; i++ {
			//mu.Lock()
			a[i] = i
			//mu.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i <= 10000; i++ {
			//mu.Lock()
			a[i] = i * 2
			//mu.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i <= 10000; i++ {
			//mu.Lock()
			a[i] = i * 3
			//mu.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i <= 10000; i++ {
			//mu.Lock()
			a[i] = i * 4
			//mu.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	res := [4]int{}
	for i, v := range a {
		//fmt.Println(i," ",v)
		if v == i {
			res[0]++
		}
		if v == i*2 {
			res[1]++
		}
		if v == i*3 {
			res[2]++
		}
		if v == i*4 {
			res[3]++
		}
	}
	fmt.Printf("goroutine1: %d\ngoroutine2: %d\ngoroutine3: %d\ngoroutine4: %d\n", res[0], res[1], res[2], res[3])
	fmt.Println(a)
}
