package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	x, y, z, u, v := 0, 0, 0, 0, 0
	for i := 1; ; {
		select {
		case ch <- i:
			if i > 10000 {
				goto label
			}
			fmt.Println("send into chan: ", i)
			i++
		case <-ch:
			x++
		case <-ch:
			y++
		case <-ch:
			z++
		case <-ch:
			u++
		case <-ch:
			v++
		}
	}
label:
	fmt.Println("end of all")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(u)
	fmt.Println(v)
}
