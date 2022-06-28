package main

import (
	"fmt"
	"math"
	"sync/atomic"
	"unsafe"
)

func atomicLoadFloat64(x *float64) float64 {
	return math.Float64frombits(atomic.LoadUint64((*uint64)(unsafe.Pointer(x))))
}

func main() {
	var a float64
	a = 55.22
	b := atomicLoadFloat64(&a)
	fmt.Println(b)

	var aba bool
	fmt.Println(aba)
}
