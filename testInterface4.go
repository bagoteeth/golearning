package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var w io.Writer
	fmt.Printf("%T\n%+v\n", w, w)
	fmt.Println(w == nil)
	w = new(bytes.Buffer)
	fmt.Printf("%T\n%+v\n", w, w)
	fmt.Println(w == nil)
	w.Write([]byte("hello"))
	fmt.Printf("%T\n%+v\n", w, w)
	x := w.(*bytes.Buffer)
	x.WriteString("world")
	fmt.Printf("%T\n%+v\n", x, x.String())
}
