package main

import (
	"fmt"
	"strings"
)

func main() {
	var a interface{}
	var b interface{}
	a = 12345678
	b = 23456789
	fmt.Println(a > b)
}
