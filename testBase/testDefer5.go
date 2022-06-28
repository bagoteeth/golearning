package main

import "fmt"

func main() {
	defer func() {
		fmt.Println(1)
		fmt.Println(2)
	}()
	fmt.Println(3)
}
