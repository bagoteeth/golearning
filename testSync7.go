package main

import (
	"fmt"
	"sync"
)

type onceStruct struct {
	name string
}

var once sync.Once
var instance *onceStruct

func getInstance() *onceStruct {
	once.Do(func() {
		instance = &onceStruct{name: "testonce"}
		fmt.Println("only run once")
	})
	return instance
}
func main() {
	a := getInstance()
	b := getInstance()
	c := getInstance()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
