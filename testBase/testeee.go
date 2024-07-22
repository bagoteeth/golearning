package main

import (
	"fmt"
)

type testeee struct {
	Name string
	Age  int
}

type Severity string

func main() {
	map1 := map[string]interface{}{"a": 1, "b": "b", "c": "abab"}
	map2 := map[string]interface{}{"a": 1, "b": 1, "c": "abaab"}
	fmt.Println(mapIntersection(map1, map2))

}

func mapIntersection(m1, m2 map[string]interface{}) bool {
	for key, _ := range m1 {

		_, ok := m2[key]
		if !ok {
			continue
		}
		if m1[key] == m2[key] {
			return true
		}
	}
	return false
}
