/*
*

	@author rfjiang
	@date 2024/7/12
	@desc

*
*/
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	var b []int
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(data, &b)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(b)
	b[2] = 1000
	fmt.Println(b)
	fmt.Println(a)
	for i := range b {
		b[i] = 100
	}
	fmt.Println(b)
	fmt.Println(a)

}
