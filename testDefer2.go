package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("running defer")
	}()
	fmt.Println("running")
}
