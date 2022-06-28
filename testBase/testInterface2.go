package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{11, 22, 33, 44}
	var b []int
	b = a
	fmt.Println(b)
	var c []interface{}
	for _, v := range a {
		c = append(c, v)
	}
	fmt.Println(c)
	var d interface{} = 15
	fmt.Println(reflect.TypeOf(d).Kind())
}
