package main

import "fmt"

type testmap struct {
	Items map[string]abc
}
type abc map[string]bool

func main() {
	//a:=testmap{
	//	Items: map[string]abc{},
	//}
	//
	//
	//fmt.Println(a.Items["a"])
	//fmt.Println(a.Items["c"])

	a := make(map[string]bool)
	a["a"] = true
	a["b"] = false
	fmt.Println(a["a"])
	fmt.Println(a["b"])
	fmt.Println(a["c"])
}
