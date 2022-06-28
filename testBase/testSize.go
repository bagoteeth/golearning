package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 50
	b := 5000000000
	fmt.Printf("%T %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("%T %d\n", b, unsafe.Sizeof(b))
	var c int32 = 5000
	var d int64 = 50000000000000
	fmt.Printf("%T %d\n", c, unsafe.Sizeof(c))
	fmt.Printf("%T %d\n", d, unsafe.Sizeof(d))

	e := "a"
	f := "afdg;oiergihrglerglahgdflbvdlafgkjrehkgelhablrdddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddde"
	fmt.Printf("%T %d\n", e, unsafe.Sizeof(e))
	fmt.Printf("%T %d\n", f, unsafe.Sizeof(f))

	g := true
	fmt.Printf("%T %d\n", g, unsafe.Sizeof(g))

	h := struct {
		name string  //16
		num  float32 //4
		num2 int32   //4
		//age int//8
		//age2 int//8
		//ha bool//1
		//haha bool//1
		//haha2 bool//1
		//haha3 bool//1
		//haha4 bool//1
		//haha5 bool//1
		//haha6 bool//1
		//haha7 bool//1
		//haha8 bool//1
	}{"bago", 1.12121121, 2}
	fmt.Printf("%T %d\n", h, unsafe.Sizeof(h))

	var i float32 = 1.12
	fmt.Printf("%T %d\n", i, unsafe.Sizeof(i))

}
