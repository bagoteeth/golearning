package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type() //把s.Type()返回的Type对象复制给typeofT，typeofT也是一个反射。
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i) //迭代s的各个域，注意每个域仍然是反射。
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface()) //提取了每个域的名字
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}
