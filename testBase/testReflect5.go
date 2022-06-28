package main

import (
	"fmt"
	"reflect"
)

type testReflect5 struct {
	Name string
	Ages ages
	Num  int8
}
type ages struct {
	A int
	B int
}

func main() {
	t := testReflect5{
		Name: "aaa",
		Ages: ages{3, 33},
		Num:  11,
	}
	e := reflect.ValueOf(&t).Elem()
	typeE := e.Type()
	for i := 0; i < e.NumField(); i++ {
		f := e.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeE.Field(i).Name, f.Type(), f.Interface())
	}
	e.Field(0).SetString("bbb")
	e.Field(1).Set(reflect.ValueOf(ages{
		A: 9,
		B: 99,
	}))
	e.Field(2).SetInt(65663) //65536+127
	fmt.Println("*************")
	for i := 0; i < e.NumField(); i++ {
		f := e.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeE.Field(i).Name, f.Type(), f.Interface())
	}
}
