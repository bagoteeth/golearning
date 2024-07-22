package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)
	fmt.Println(len(m))
	m[1]++
	m[2]++
	m[2]++
	fmt.Println(len(m))
	fmt.Println(m)
	m[2]--
	if m[2] == 0 {
		delete(m, 2)
	}
	fmt.Println(len(m))
	fmt.Println(m)
	m[2]--
	if m[2] == 0 {
		delete(m, 2)
	}
	fmt.Println(len(m))
	fmt.Println(m)
}
