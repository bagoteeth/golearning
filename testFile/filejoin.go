/*
*

	@author rfjiang
	@date 2024/6/24
	@desc

*
*/
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	a := "/home/ubt22-dt-24-01/Oracle/Middleware/Oracle_Home/wlserver/samples/server/docs/WEB-INF/xxx.xml"
	b := "../examples/build"
	c := filepath.Join(a, b)
	fmt.Println(c)
}
