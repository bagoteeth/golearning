/*
*

	@author rfjiang
	@date 2024/7/10
	@desc

*
*/
package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name string  `xml:"name"`
	Age  *int    `xml:"age"`
	City *string `xml:"city"`
}

func main() {
	xmlData := []byte(`
		<person>
			<name>Alice</name>
			<age></age>
		</person>
	`)
	xmlData2 := []byte(`
		<person>
			<name>Alice</name>
			<city></city>
		</person>
	`)
	xmlData3 := []byte(`
		<person>
			<name>Alice</name>
			<city></city>
			<rfjiang>abc</rfjiang>
		</person>
	`)

	var p1, p2, p3 Person
	err := xml.Unmarshal(xmlData, &p1)
	if err != nil {
		fmt.Printf("err1: %s\n", err.Error())
	}
	err = xml.Unmarshal(xmlData2, &p2)
	if err != nil {
		fmt.Printf("err2: %s\n", err.Error())
	}
	err = xml.Unmarshal(xmlData3, &p3)
	if err != nil {
		fmt.Printf("err3: %s\n", err.Error())
	}

	fmt.Printf("p1:\n%+v\np2:\n%+v\np3:\n%+v\n", p1, p2, p3)

	fmt.Println(p1.City)
	fmt.Println(p2.City)
}
