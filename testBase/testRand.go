package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(strconv.Itoa(rand.Intn(50)))
	fmt.Println(strconv.Itoa(rand.Intn(50)))
	fmt.Println(strconv.Itoa(rand.Intn(50)))
	fmt.Println(strconv.Itoa(rand.Intn(50)))
	fmt.Println(rand.Intn(50))
	fmt.Println(rand.Intn(50))
	fmt.Println(rand.Intn(50))
	fmt.Println(rand.Intn(50))
}
