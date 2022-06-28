package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer")
	}()
	fmt.Println(1)
	panic("panicing")
	fmt.Println(2)
}
