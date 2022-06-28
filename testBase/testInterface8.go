package main

import (
	"fmt"
)

func main() {
	var a interface{} = 5
	b := a.(int)
	fmt.Println(b + 2)
}
