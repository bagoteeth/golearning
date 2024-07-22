/*
*

	@author rfjiang
	@date 2024/5/22
	@desc

*
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	//ch1Addr := reflect.ValueOf(ch1).Pointer()
	//fmt.Println(unsafe.Pointer(ch1Addr))
	//go func() {
	//	for true {
	//		a := <- ch1
	//		fmt.Printf("ch1 receive: %d\n", a)
	//		time.Sleep(3 * time.Second)
	//	}
	//}()
	//go func() {
	//	for true {
	//		a := <- ch2
	//		fmt.Printf("ch2 receive: %d\n", a)
	//		time.Sleep(5 * time.Second)
	//	}
	//}()
	//go func() {
	//	for true {
	//		a := <- ch3
	//		fmt.Printf("ch3 receive: %d\n", a)
	//		time.Sleep(7 * time.Second)
	//	}
	//}()
	//
	//
	//for i := 0;;i++ {
	//	select {
	//	case ch1 <- i:
	//		fmt.Printf("ch1 send: %d\n", i)
	//	case ch2 <- i:
	//		fmt.Printf("ch2 send: %d\n", i)
	//	case ch3 <- i:
	//		fmt.Printf("ch3 send: %d\n", i)
	//
	//	}
	//	time.Sleep(time.Second)
	//}

dflt:
	fmt.Println("dflt")

case1:
	fmt.Println("case1")

case2:
	fmt.Println("case2")

case3:
	fmt.Println("case3")

	var chosen int
	var recvOK bool

	select {
	case ch1 <- 1:
		goto case1
	case ch2 <- 1:
		goto case2
	case ch3 <- 1:
		goto case3
	default:
		goto dflt
	}

	//if chosen < 0 {
	//	goto dflt
	//}
	//if chosen == 0 {
	//	goto case1
	//}
	//if chosen == 1 {
	//	goto case2
	//}
	//goto case3

dflt:
	fmt.Println("dflt")

case1:
	fmt.Println("case1")

case2:
	fmt.Println("case2")

case3:
	fmt.Println("case3")

}
