package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(add))
	fmt.Println(reflect.TypeOf(minus))
	test(3, 4, add)
	test(7, 9, minus)
}
func add(a int, b int) int {
	return a + b
}
func minus(a int, b int) int {
	return a - b
}
func test(x int, y int, f func(int, int) int) {
	fmt.Println(f(x, y))
}
