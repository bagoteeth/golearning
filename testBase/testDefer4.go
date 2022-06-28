package main

import "fmt"

func main() {
	fmt.Println(testDefer())
}

func testDefer() (r int) {
	temp := 10
	defer func(r int) {
		r++
	}(r)
	return temp
}
