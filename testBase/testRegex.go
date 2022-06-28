package main

import (
	"fmt"
	"regexp"
)

func main() {
	a, b := regexp.MatchString(`^\w*$`, "_2__sgh64fs__fgh65a")
	fmt.Println(a)
	fmt.Println(b)
}
