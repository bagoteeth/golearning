package main

import (
	"fmt"
	"time"
)

func main() {
	testTimer()
}

func testTimer() {
	defer doTimer(time.Now())
	time.Sleep(5 * time.Second)
}

func doTimer(start time.Time) {
	fmt.Println(start)
	fmt.Println(time.Now())
	fmt.Println(time.Since(start))
}
