package main

import (
	"fmt"
)

type duck interface {
	move()
	getName() string
}
type bird struct {
	name string
	age  int
}

func (b *bird) move() {
	fmt.Println("flying")
}
func (b *bird) getName() string {
	return b.name
}
func main() {
	var d duck = &bird{
		name: "bird1",
	}
	d.move()
	fmt.Println(d.getName())
	fmt.Printf("%T", d)
}
