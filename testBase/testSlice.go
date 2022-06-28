package main

import "fmt"

func main() {

	var a []int
	for i := 1; i <= 2000; i++ {
		a = append(a, i)
		fmt.Printf("%p\n", &a)

	}
	b := make(map[string]int)
	b["aaa"] = 1
	b["bbb"] = 2
	fmt.Printf("%p\n", &b)
}
