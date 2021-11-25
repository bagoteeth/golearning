package main

import "fmt"

type sStruct struct {
	name string
	Age  int
}

func main() {
	s := sStruct{
		name: "aaaa",
		Age:  5,
	}
	s.Age = 6
	s.name = "bbbb"
	fmt.Println(s)
}
