package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = "ababa"
	t := reflect.TypeOf(a)
	fmt.Println(t)
}
