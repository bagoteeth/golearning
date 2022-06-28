package main

import (
	"fmt"
	"reflect"
)

type otherint int
type D struct {
	name string
	age  int
}

func main() {
	var a otherint = 5
	t := reflect.TypeOf(a)
	fmt.Println(t)
	fmt.Println(t.Name())
	fmt.Println(t.Kind())
	var b int8 = 33
	var c float64 = 44.4
	vb := reflect.ValueOf(b)
	vc := reflect.ValueOf(c)
	vb.Elem()
	fmt.Println(vb)
	fmt.Println(vc)
	//b := struct {
	//	a int
	//	b string
	//}{1, "abab"}
	//t2 := reflect.TypeOf(b)
	//fmt.Println(t2)
	//fmt.Println(t2.Name())
	//fmt.Println(t2.Kind())
	//c := "adfg"
	//t3 := reflect.TypeOf(c)
	//fmt.Println(t3)
	//fmt.Println(t3.Name())
	//fmt.Println(t3.Kind())
	//d := D{
	//	Name: "ddddddd",
	//	age:  24,
	//}
	//t4 := reflect.TypeOf(d)
	//fmt.Println(t4)
	//fmt.Println(t4.Name())
	//fmt.Println(t4.Kind())
	//var e rune = 65
	//t5 := reflect.TypeOf(e)
	//fmt.Println(t5)
	//fmt.Println(t5.Name())
	//fmt.Println(t5.Kind())
}
