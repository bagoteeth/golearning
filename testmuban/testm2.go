/*
*

	@author rfjiang
	@date 2024/5/21
	@desc

*
*/
package testmuban

// S1
type S1 struct {
	aaa string
	bbb int
}

// name
type name interface {
	F1()
	//  @title       F2
	//  @description
	//  @param       a
	//  @param       b
	//  @return      int
	F2(a int, b string) int
}

// @title       F11
// @description
// @param       a
// @param       b
// @param       c
// @return      int
// @return      error
func F11(a int, b string, c S1) (int, error) {
	return 0, nil
}

// @title       F22
// @description
// @receiver    i
// @param       a
// @param       b
// @return      res
// @return      err
func (i *S1) F22(a int, b string) (res string, err error) {
	return "", nil
}
