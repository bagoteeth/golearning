package main

//
//import (
//	"fmt"
//	"math"
//)
//
//type shape interface {
//	area() float64
//}
//type circle struct {
//	radius float64
//}
//
//func (c *circle) area() float64 {
//	return math.Pi * c.radius * c.radius
//}
//
//type rectangle struct {
//	length float64
//	width  float64
//}
//
//func (r *rectangle) area() float64 {
//	return r.length * r.width
//}
//func calculateArea(s shape) {
//	fmt.Printf("area of %T is %f\n", s, s.area())
//}
//func main() {
//	mycircle := &circle{
//		radius: 5,
//	}
//	myrectangle := &rectangle{
//		length: 4,
//		width:  6,
//	}
//	calculateArea(mycircle)
//	calculateArea(myrectangle)
//}
