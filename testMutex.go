package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Done()
	wg.Done()

	fmt.Println("done")
}
