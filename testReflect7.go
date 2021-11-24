package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = 1
	v := reflect.TypeOf(&a)
	fmt.Println(v)
	fmt.Println(v.Name())
	fmt.Println(v.Kind())
	e := v.Elem()
	fmt.Println(e.Name())
	fmt.Println(e.Kind())
}
