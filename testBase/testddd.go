package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "2@rfjiang111@default@worker@@@"
	s2 := "3@rfjiang111@@@kube-system@DaemonSet@kube-flannel-ds"
	fmt.Println(strings.Split(s1, "@"))
	fmt.Println(strings.Split(s2, "@"))

	fmt.Println(len(strings.Split(s1, "@")))
	fmt.Println(len(strings.Split(s2, "@")))

}
