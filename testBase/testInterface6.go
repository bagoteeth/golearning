package main

import (
	"fmt"
	"strings"
)

func main() {
	var a interface{} = "hello world"
	if a, ok := a.(string); ok {
		fmt.Println(strings.ToUpper(a))
	}
}
