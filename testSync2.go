package main

import (
	"fmt"
	"time"
)

var share_cnt uint64 = 0

func incrShareCnt() {
	for i := 0; i < 1000000; i++ {
		share_cnt++
	}

	fmt.Println(share_cnt)
}

func main() {

	for i := 0; i < 2; i++ {
		go incrShareCnt()
	}

	time.Sleep(5 * time.Second)

}
