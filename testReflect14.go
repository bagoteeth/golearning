package main

import (
	"fmt"
	"reflect"
)

type testR14 struct {
	Name string
	Age  int
}

func main() {
	a := testR14{
		Name: "ababab",
		Age:  7,
	}
	va := reflect.ValueOf(&a).Elem()
	fmt.Println(va.CanSet()) //CanSet() bool	返回值能否被修改。要求值可寻址且是导出的字段
	va.Field(0).SetString("ccccccc")
	va.Field(1).SetInt(9)
	fmt.Println(a)
}
