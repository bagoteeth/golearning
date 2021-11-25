package main

import (
	"fmt"
	"reflect"
)

type testR12 struct {
	Name string
	Age  int
}

func main() {
	var a int8 = 10
	b := "abc"
	c := testR12{
		Name: "ccc",
		Age:  22,
	}
	va := reflect.ValueOf(a) //Value变量
	vb := reflect.ValueOf(b)
	vc := reflect.ValueOf(c)
	fmt.Printf("type: %T, value: %d\n", va.Int(), va.Int()) //原始值
	fmt.Printf("type: %T, value: %s\n", vb.String(), vb.String())
	fmt.Printf("type: %T, value: %+v\n", vc.Interface().(testR12), vc.Interface().(testR12))
}
