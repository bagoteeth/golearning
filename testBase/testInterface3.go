package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer = os.Stdout
	w.Write([]byte("hello"))

	var a interface{} = []int{1, 2, 3}
	var b interface{} = []int{3, 4, 5}
	fmt.Println(a == b)
}
