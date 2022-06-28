package main

import (
	"fmt"
	"os"
)

func main() {
	processname := os.Args[0]
	fmt.Println(processname)
}
