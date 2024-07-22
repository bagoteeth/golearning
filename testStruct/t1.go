package main

import (
	"encoding/json"
	"fmt"
)

type Son struct {
	Name string `json:"nameJson"`
	Age  int    `json:"ageJson"`
}

type Fa struct {
	FName string `json:"f_nameJson"`
	Son
}

func main() {
	a := []byte(`{"f_nameJson":"bago","nameJson":"teeth","ageJson":3}`)
	var b Fa
	err := json.Unmarshal(a, &b)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", b)
	c, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", string(c))
}
