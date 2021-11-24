package main

import (
	"fmt"
	"sync"
)

var rwSync sync.Mutex

func ff() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func g() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func h() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
func main() {
	rwSync.Lock()
	defer rwSync.Unlock()
	fmt.Println(ff())
	fmt.Println(g())
	fmt.Println(h())
}
