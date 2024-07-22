package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m["a"] = append(m["a"], "a")
	m["b"] = append(m["b"], "a")
	m["b"] = append(m["b"], "b")
	fmt.Println(m)
	fmt.Println(m["a"])
	fmt.Println(m["b"])
	fmt.Println(m["c"])

}
