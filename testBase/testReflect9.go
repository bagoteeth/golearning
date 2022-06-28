package main

import (
	"fmt"
	"reflect"
)

type reflect9 struct {
	name string `mytag:"myname" othertag:"othername"` //格式为：`key1:"value1" key2:"value2"
	age  int    `mytag:"myage"`
}

func main() {

	r := reflect9{
		name: "aaa",
		age:  1,
	}
	e := reflect.TypeOf(r)
	for i := 0; i < e.NumField(); i++ {
		fmt.Printf("%d %v %v\n", i, e.Field(i).Tag.Get("mytag"), e.Field(i).Tag.Get("othertag"))
	}

}
