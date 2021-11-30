package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	m := make(map[string]int)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(4)
	go func() {
		for i := 1; i <= 10000; i++ {
			mu.Lock()
			m[strconv.Itoa(i)] = i
			mu.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i <= 10000; i++ {
			mu.Lock()
			m[strconv.Itoa(i)] = i * 2
			mu.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i <= 10000; i++ {
			mu.Lock()
			m[strconv.Itoa(i)] = i * 3
			mu.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i <= 10000; i++ {
			mu.Lock()
			m[strconv.Itoa(i)] = i * 4
			mu.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	res := [4]int{}
	for i, v := range m {
		key, _ := strconv.Atoi(i)
		if v == key {
			res[0]++
		}
		if v == key*2 {
			res[1]++
		}
		if v == key*3 {
			res[2]++
		}
		if v == key*4 {
			res[3]++
		}
	}
	fmt.Printf("goroutine1: %d\ngoroutine2: %d\ngoroutine3: %d\ngoroutine4: %d\n", res[0], res[1], res[2], res[3])
}
