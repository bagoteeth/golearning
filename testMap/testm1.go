/*
*

	@author rfjiang
	@date 2024/5/29
	@desc

*
*/
package main

import "fmt"

func main() {
	m := make(map[string]int)
	f1(&m)
	fmt.Println(m)
}

func f1(m *map[string]int) {
	(*m)["a"] = 1
}
