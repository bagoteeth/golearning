package main

import (
	"encoding/json"
	"fmt"
)

type reflect8 struct {
	Name   string `json:"myname"`
	Age    int    `json:"myage,omitempty"`
	Ignore string `json:"-"`
}

func main() {
	r1 := &reflect8{Name: "name1", Age: 1, Ignore: "aaa"}
	r2 := &reflect8{Name: "name2", Ignore: "bbb"}
	r3 := &reflect8{Age: 3, Ignore: "ccc"}
	j1, _ := json.Marshal(r1) //序列化
	j2, _ := json.Marshal(r2)
	j3, _ := json.Marshal(r3)
	fmt.Println("json1: ", string(j1))
	fmt.Println("json2: ", string(j2))
	fmt.Println("json3: ", string(j3))
	recv1 := reflect8{}
	recv2 := reflect8{}
	recv3 := reflect8{}
	json.Unmarshal(j1, &recv1) //反序列化
	json.Unmarshal(j2, &recv2)
	json.Unmarshal(j3, &recv3)
	fmt.Printf("recv1: %+v\n", recv1)
	fmt.Printf("recv2: %+v\n", recv2)
	fmt.Printf("recv3: %+v\n", recv3)
}
