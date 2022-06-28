package main

import (
	"fmt"
	"reflect"
)

type testR13 struct {
	Name string
	Age  int8
}

func main() {
	t := testR13{
		Name: "abc",
		Age:  33,
	}
	v := reflect.ValueOf(t)
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("type: %s value: %v\n", v.Field(i).Type().Name(), v.Field(i))
	}
}
