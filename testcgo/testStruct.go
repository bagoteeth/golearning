package main

/*
struct A {
	int type;
};
*/
import "C"
import "fmt"

func main() {
	var a C.struct_A
	fmt.Println(a._type)
}
