package main

import (
	"fmt"
)

func main() {
	sum, product := Calculate(5, 7)
	fmt.Println(sum)
	fmt.Println(product)
}
func Calculate(a int, b int) (c int, d int) {
	c = a + b
	d = a - b
	return
}
