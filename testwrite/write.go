package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Create("data.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()

	for i := 0; i < 10; i++ {
		_, err = f.WriteString(strconv.Itoa(i) + "\n")
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}
