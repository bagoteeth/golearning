package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := BigFilter{
		FA: []FilterA{
			{Name: "aaa"},
			{Name: "bbb"},
			{Name: "cac"},
		},
		FB: []FilterB{
			{Name: "aaa"},
			{Name: "bbb"},
			{Name: "bcb"},
		},
	}
	result := BigFilter{}

	refA := reflect.ValueOf(&a).Elem()
	refR := reflect.ValueOf(&result).Elem()
	field := refA.NumField()
	for i := 0; i < field; i++ {
		if refA.Field(i).Kind() == reflect.Slice {
			for j := 0; j < refA.Field(i).Len(); j++ {
				vintf := refA.Field(i).Index(j).Addr().Interface().(FilterIntf)
				if vintf.A() {
					newS := reflect.Append(refR.Field(i), refA.Field(i).Index(j))
					refR.Field(i).Set(newS)
				}
			}
		}
	}
	fmt.Println(result)
}
