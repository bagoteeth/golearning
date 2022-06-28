package main

import (
	"encoding/json"
	"fmt"
)

type testStruct2 struct {
	Name string `json:"Name"`
	Fack struct {
		Prom  string `json:"prom"`
		Grafa string `json:"grafa"`
	} `json:"fack"`
}

type testStruct22 struct {
	Name string     `json:"Name"`
	Fack FackStruct `json:"fack"`
}

type FackStruct struct {
	Sheet string `json:"sheet,omitempty"`
	Freak bool   `json:"freak"`
}

func (t *testStruct22) add(a, b, c string) {
	t.Name = a

	t.Fack.Sheet = c
}

func main() {
	//a:= &testStruct22{}
	//a.add("absadb","asdfbdafb","qweqwweq")
	//fmt.Printf("%+v\n",*a)

	b := FackStruct{
		Sheet: "abas",
	}
	bb, _ := json.Marshal(b)
	bbb := string(bb)
	fmt.Printf("%+v\n", bbb)

	c := FackStruct{
		Sheet: "iuoiouio",
		Freak: false,
	}
	cc, _ := json.Marshal(c)
	ccc := string(cc)
	fmt.Printf("%+v\n", ccc)
	//jsonstr:=`{"fack":{"prom":"jp","grafa":"jg"}}`
	//t:=testStruct2{
	//	Name: "asdsad",
	//	Fack: struct {
	//		Prom  string `json:"prom"`
	//		Grafa string `json:"grafa"`
	//	}{
	//		Prom: "Afgfdb",
	//		Grafa: "gnhgmgh",
	//	},
	//}
	//fmt.Println(t)
	//
	//
	//u:= testStruct22{
	//	Name: "hhhhh",
	//	Fack: FackStruct{
	//		Sheet: "iii",
	//		Freak: "jjjjj",
	//	},
	//}
	//fmt.Println(u)
	//
	//c:=testStruct2{}
	//json.Unmarshal([]byte(jsonstr),&c)
	//fmt.Printf("%+v\n",c)
}
