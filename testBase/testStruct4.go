package main

import (
	"encoding/json"
	"fmt"
)

type testStruct4 struct {
	Id string
	T  testStruct44
}

type testStruct44 struct {
	Name string
}

func main() {
	//a:=testStruct4{
	//	Id: "bago",
	//	T: testStruct44{Name: "teeth"},
	//}
	//b:=testStruct4{
	//	Id: "bago",
	//	T: testStruct44{Name: "teeth"},
	//}
	//fmt.Println(a==b)

	var j interface{}
	j = `[
    {
        "Name": "aaa",
        "id": "bbb"
    },
    {
        "Name": "aaa",
        "id": "bbb"
    },
    {
        "Name": "ccc",
        "id": "ddd"
    }
]`
	jj, _ := json.Marshal(j)
	fmt.Println(string(jj))
	var jjj []map[string]string
	err := json.Unmarshal([]byte(string(jj)), &jjj)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(jjj)
}
