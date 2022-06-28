package main

import (
	"fmt"
	"reflect"
)

func main() {
	map1 := make(map[string]int)
	map2 := make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2
	map1["c"] = 3
	map2["c"] = 3
	map2["b"] = 2
	map2["a"] = 1
	fmt.Println(reflect.DeepEqual(map1, map2))
}
