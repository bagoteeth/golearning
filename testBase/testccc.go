package main

import (
	"fmt"
	"reflect"
)

type Testccc struct {
	A []Testccc1
	B []Testccc1
}

type Testccc2 interface {
}

type Testccc1 struct {
	Name string
}

func main() {
	a1 := Testccc1{Name: "aaa"}
	a2 := Testccc1{Name: "bbb"}
	a3 := Testccc1{Name: "ccc"}

	b := Testccc{
		A: []Testccc1{a1, a2},
		B: []Testccc1{a1, a3},
	}
	fmt.Println(isin(a2, []Testccc1{a1, a2}))
	fmt.Println(isin(a2, b.B))
}

func isin(a, as interface{}) bool {
	ass := reflect.ValueOf(as)
	for i := 0; i < ass.Len(); i++ {
		if a == ass.Index(i).Interface() {
			return true
		}
	}
	return false
}
