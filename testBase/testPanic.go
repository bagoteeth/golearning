package main

import "fmt"

func testPanic() {
	panic("发生恐慌了")
}
func testPanic2() {
	panic("发生恐慌了2")
}
func division(x, y int) (result int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	result = x / y
	return result, nil
}
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("异常信息为：%+v", err)
		}
	}()
	//defer func() {
	//	fmt.Println(1)
	//}()
	//defer func() {
	//	fmt.Println(2)
	//}()
	//a, err := division(5, 0)
	//fmt.Printf("%d\n%+v\n", a, err)
	fmt.Println(1)
	testPanic()
	fmt.Println(2)
	//testPanic2()
}
