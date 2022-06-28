package main

import (
	"encoding/json"
	"fmt"
)

type testStruct5 struct {
	Name  testStruct55 `json:"name"`
	Name2 testStruct55 `json:"name2"`
}

type testStruct55 struct {
	Val string `json:"val"`
}

func main() {
	a := `
{name:{val:"aaa"}}, name2:{val:"bbb"}}
`
	b := testStruct5{
		Name:  testStruct55{Val: "ffffff"},
		Name2: testStruct55{Val: "sssssss"},
	}
	var mmp testStruct5
	aJson, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", string(aJson))
	err = json.Unmarshal(aJson, &mmp)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(mmp)
	fmt.Printf("%+v\n", b)
	//	var mmp interface{}
	//	var mmp2 interface{}
	//	mmp = testStruct5{
	//		Name:  testStruct55{Val: "aaa"},
	//		Name2: testStruct55{Val: "bbb"},
	//	}
	//	mmp2 = testStruct5{
	//		Name:  testStruct55{Val: "aaa"},
	//		Name2: testStruct55{Val: "bbb"},
	//	}
	//	fmt.Println(mmp==mmp2)
}
