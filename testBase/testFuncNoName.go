package main

import "fmt"

func main() {
	a := 5
	for i := 0; i < 3; i++ {
		func() {
			a++
			fmt.Println(a)
		}()
	}
}
