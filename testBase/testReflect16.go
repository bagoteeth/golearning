package main

import (
	"fmt"
	"reflect"
)

type mystruct16 struct {
	Name string
}

func main() {
	ms := mystruct16{Name: "ababab"}
	fmt.Printf("%+v\n", reflect.TypeOf(ms))
	fmt.Printf("%+v\n", reflect.TypeOf(ms).Name())
	a := reflect.TypeOf(ms).Kind()
	fmt.Printf("%T	%+v\n", a, a)
}
