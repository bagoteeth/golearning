package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	unix := now.Unix()
	fmt.Println(unix)
	utcTime := now.UTC()
	fmt.Println(utcTime)
	timeLayout := "2006-01-02 15:04:05"
	// time.Unix的第二个参数传递0或10结果一样，因为都不大于1e9
	timeStr := time.Unix(unix, 0).Format(timeLayout)
	fmt.Println(timeStr)
}
