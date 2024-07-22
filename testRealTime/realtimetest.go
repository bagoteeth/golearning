package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

const (
	SCHED_NORMAL = iota
	SCHED_FIFO
	SCHED_RR
	SCHED_BATCH
	SCHED_ISO
	SCHED_IDLE
	SCHED_DEADLINE
)

func main() {
	pid := syscall.Getpid()
	fmt.Println("Process id is", pid)

	fmt.Println("Scheduler vorher:")
	syscall.Syscall(145, uintptr(0), 0, 0)

	var schedParam uint32 = 9
	var schedType = SCHED_FIFO
	_, _, err := syscall.Syscall(uintptr(144), uintptr(0), uintptr(schedType), uintptr(unsafe.Pointer(&schedParam)))
	fmt.Println("Scheduler setzen", err)

	fmt.Println("Scheduler nachher:")
	syscall.Syscall(145, uintptr(0), 0, 0)

	for {
		fmt.Println("Realtime task is running...")
		time.Sleep(5 * time.Second)
	}
}
