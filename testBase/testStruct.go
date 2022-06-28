package main

import "fmt"

type sStruct struct {
	name *string
	Age  *int
}

func (s *sStruct) Dosth() {
	fmt.Printf("fuck %s you %d\n", s.name, s.Age)
}
func main() {
	mmp := make(map[string]*sStruct)
	var a sStruct
	b := sStruct{}
	var c sStruct = sStruct{}
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", c)

	a.Dosth()
	b.Dosth()
	c.Dosth()
	mmp["a"] = &a
	mmp["b"] = &b
	mmp["c"] = &c
	mmp["a"].Dosth()
	mmp["d"] = &sStruct{
		name: nil,
		Age:  nil,
	}
	fmt.Println(mmp["d"])
}
