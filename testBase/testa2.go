package main

import (
	"fmt"
	"strings"
)

type Fack1 struct {
	name string
}

var fack1map map[string]*Fack1

func main() {
	f := "  asdaf   sadfasdfdf sadfasdfsda   sadfsads s sd "
	fmt.Println(strings.TrimSpace(f))
}

func Fack1do() {
	f := Fack1{name: "sheet"}
	fack1map["a"] = &f
}
