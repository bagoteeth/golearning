package main

import (
	"fmt"
	"net"
)

func main() {
	nets, err := net.Interfaces()
	if err != nil {
		fmt.Printf("err: %s", err.Error())
	}
	fmt.Println(nets)
}
