package main

import "fmt"

func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum
}
func main() {
	s := []int{1, 3, 5, 7, 9, 2, -4}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	//go sum(s,c)
	go sum(s[len(s)/2:], c)
	x := <-c
	y := <-c
	//z:=<-c
	fmt.Println(x)
	fmt.Println(y)
	//fmt.Println(z)
}
