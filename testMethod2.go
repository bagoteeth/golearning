package main

import "fmt"

type Mymap map[string]int

func (r Mymap) f() {
	r["test"] = 100
}
func main() {
	a := Mymap{}
	a.f()
	fmt.Printf("%+v", a)
}
