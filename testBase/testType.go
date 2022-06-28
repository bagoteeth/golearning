package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 5
	fmt.Println(reflect.TypeOf(&a).Kind() == reflect.Ptr)

	var b uint16 = 555
	fmt.Println(reflect.TypeOf(b).Kind() >= reflect.Int && reflect.TypeOf(b).Kind() <= reflect.Float64)
}
