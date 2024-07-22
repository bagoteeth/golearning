package main

import (
	"fmt"
	"time"
)

func main() {
	//now := time.Now()
	//fmt.Println(now)
	//unix := now.Unix()
	//fmt.Println(unix)
	//utcTime := now.UTC()
	//fmt.Println(utcTime)
	//timeLayout := "2006-01-02 15:04:05"
	//// time.Unix的第二个参数传递0或10结果一样，因为都不大于1e9
	//timeStr := time.Unix(unix, 0).Format(timeLayout)
	//fmt.Println(timeStr)
	//time.Sleep(3*time.Second)
	//latenow := time.Now()
	//fmt.Println(now.After(latenow))
	//fmt.Println(now.Before(latenow))
	//fmt.Println(now.Before(now))
	//fmt.Println(now.After(now))
	//
	//var timeinterface interface{}
	//timeinterface = time.Now()
	//fmt.Println(timeinterface.(time.Time))
	//
	//a:= 1.23
	//fmt.Println(int(a))
	//a:="2022-07-01T18:20:09+08:00"
	//b, _ := time.Parse(time.RFC3339,a)
	//c, _ := time.Parse(time.RFC3339,a)
	//
	//fmt.Printf("%+v\n",b)
	//fmt.Println(reflect.TypeOf(b))
	//fmt.Println(b==c)

	//fmt.Printf(`aaaa'%%%s%%'`,"bbbbb")
	var b int32
	b = 2200000000
	fmt.Println(time.Now().Unix())
	a := `asdasdasd\n%d\nasdasdsd\n`
	fmt.Printf(a, b)
}
