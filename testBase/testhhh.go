package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Testhhh struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	a := make(map[Testhhh]Testhhh)
	b := Testhhh{
		Name: "abab",
		Age:  123,
	}
	c := Testhhh{
		Name: "afdgasdf",
		Age:  345345,
	}
	a[b] = Testhhh{
		Name: "ab",
		Age:  0,
	}
	fmt.Println(a[b].Name)
	fmt.Println(a[c].Age)

	var d interface{}
	d = nil
	fmt.Println(reflect.TypeOf(d))
	d = "abc"
	fmt.Println(reflect.TypeOf(d))

	e := `
asda${abab}354fsd53sd ${bcbc}55555${abab}asdasd${bcbc}
`
	e = strings.ReplaceAll(e, "${abab}", "AAAAA")

	h := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range h {
		if v%3 == 0 {
			h[i] = h[len(h)-1]
			h = append(h[:i], h[i+1:]...)

		}
	}
	fmt.Println(h)
}

func srmv(s interface{}, i int) []interface{} {
	s1 := s.([]interface{})
	return append(s1[:i], s1[i+1:]...)
}
