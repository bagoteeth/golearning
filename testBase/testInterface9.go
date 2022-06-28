package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := f()
	fmt.Println(reflect.TypeOf(a).Kind())
	p(a)

}

func f() []map[string]interface{} {
	return []map[string]interface{}{}
}
func p(a interface{}) {
	fmt.Println(a)
}
