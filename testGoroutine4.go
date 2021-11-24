package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
			fmt.Println(i)

			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func() {
		for _, i := range []string{"a", "b", "c", "d", "e"} {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	//time.Sleep(1 * time.Hour)
}
