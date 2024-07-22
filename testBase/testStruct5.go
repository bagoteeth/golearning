package main

import (
	"fmt"
)

type testStruct5 struct {
	Name  testStruct55 `json:"name"`
	Name2 testStruct55 `json:"name2"`
}

type testStruct55 struct {
	Val string `json:"val"`
}

func main() {
	a := []testStruct5{}
	fmt.Println(a == nil)
}
