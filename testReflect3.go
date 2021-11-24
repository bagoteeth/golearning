package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(5.2)
	fmt.Println(v)
	y := v.Interface().(float64) // y will have type float64.
	fmt.Println(y)
}
