package main

import "fmt"

func main() {
	fmt.Println(testRecover())
}
func testRecover() (r int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			r++
		}
	}()
	panic(struct {
		name string
		age  int
	}{"bago", 123})
	r += 5
	return
}
