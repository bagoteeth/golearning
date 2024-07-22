/*
*

	@author rfjiang
	@date 2024/6/6
	@desc

*
*/
package main

import "fmt"

type Slice1 struct {
	name string
	ages []int
}

func main() {
	res := make([]Slice1, 0)
	s0 := Slice1{
		name: "aaa",
		ages: []int{1, 2, 3},
	}
	res = append(res, s0)
	for i := 0; i < 10; i++ {
		tmp := changes(s0, i)
		res = append(res, tmp)
		//s0.ages = append(s0.ages, 10)
		//res = append(res, s0)
	}
	fmt.Println(res)
}

func changes(slice1 Slice1, i int) Slice1 {
	res := slice1
	res.ages[0] = i
	return res
}
