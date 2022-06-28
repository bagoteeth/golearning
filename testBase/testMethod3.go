package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func (p *Point) Distance(q Point) int {
	return int(math.Abs(float64(p.x-q.x)) + math.Abs(float64(p.y-q.y)))
}
func main() {
	p := Point{1, 1}
	fmt.Println(p.Distance(Point{3, 4})) //直接使用方法
	f := p.Distance                      //方法变量
	fmt.Println(f(Point{-1, 7}))
	f2 := (*Point).Distance //方法表达式
	fmt.Println(f2(&p, Point{7, 9}))
}
