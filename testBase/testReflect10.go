package main

import (
	"fmt"
	"reflect"
)

type typeStruct struct {
	name string
	age  int
}

func main() {
	t := typeStruct{name: "aaa", age: 22}
	typeT := reflect.TypeOf(t)
	for i := 0; i < typeT.NumField(); i++ {
		fmt.Printf("i: %d type: %v Name: %s\n",
			i, typeT.Field(i).Type, typeT.Field(i).Name)
	}
}
