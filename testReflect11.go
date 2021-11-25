package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{}
	t := reflect.TypeOf(a)
	fmt.Println(t)
}
