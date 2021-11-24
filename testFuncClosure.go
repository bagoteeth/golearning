package main

import "fmt"

func main() {
	f := testClosure()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
func testClosure() func() int {
	private := 10
	return func() int {
		private++
		return private
	}
}
