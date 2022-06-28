package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type money struct {
	amount int
	sync.RWMutex
}

func (m *money) deposit(amount int) {
	m.amount += amount
}

func main() {
	mymoney := money{}
	var wg sync.WaitGroup
	wg.Add(1000000)
	i := 1
	sum := 0
	t1 := time.Now()
	for ; i <= 1000; i++ {
		go func(i int) {
			mymoney.Lock()
			defer mymoney.Unlock()
			if i%2 == 0 {
				mymoney.deposit(i)
			} else {
				mymoney.deposit(-i)
			}
			wg.Done()
		}(i)
		for j := 1; j <= 999; j++ {
			go func() {
				mymoney.RLock()
				defer mymoney.RUnlock()
				for k := 1; k <= int(math.Abs(float64(mymoney.amount))); k++ {
					sum += 1
				}
				wg.Done()
			}()
		}
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(t1))
	//fmt.Println(sum)
	//fmt.Println("i: ", i)
	//fmt.Println(mymoney.amount)
}
