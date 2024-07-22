package main

import "fmt"

type intf111 interface {
	dointf()
}

type intfexm struct {
	Name string
}

func (i *intfexm) dointf() {
	fmt.Println(i.Name)
}

func main() {
	var a intf111
	a = &intfexm{Name: "abv"}
	fmt.Printf("%+v\n", a.(intfexm))
}

func map2str(m map[string]struct{}) string {
	var r string
	for key := range m {
		r = r + fmt.Sprintf("\"%s\", ", key)
	}
	if r != "" {
		r = r[:len([]rune(r))-2]
	}
	return r
}
