package main

import (
	"fmt"
)

func main() {
	a := 0
	go func() {
		a += 3
		fmt.Println("+3:", a)
		a += 4
		fmt.Println("+4:", a)
		a += 5
		fmt.Println("+5:", a)
	}()
	go func() {
		a -= 5
		fmt.Println("-5:", a)
		a -= 6
		fmt.Println("-6:", a)
		a -= 7
		fmt.Println("-7:", a)
	}()
	fmt.Println(a)
}
