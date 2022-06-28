package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type color interface {
	setColor(color string)
}

type circle struct {
	radius float64
	color  string
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c *circle) setColor(color string) {
	c.color = color
	fmt.Println("color of circle has been set to", color)
}

func main() {
	mycircle := &circle{
		radius: 5,
		color:  "white",
	}
	var myshape shape = mycircle
	var mycolor color = mycircle
	fmt.Println(myshape.area())
	mycolor.setColor("blue")
	fmt.Println(mycircle)
}
