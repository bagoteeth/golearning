package main

import (
	"fmt"
	"reflect"
)

type typeStruct2 struct {
	name string
}

func main() {
	t := &typeStruct2{name: "aaa"}
	fmt.Printf("type %s kind %s\n", reflect.TypeOf(t).Name(), reflect.TypeOf(t).Kind())
}
