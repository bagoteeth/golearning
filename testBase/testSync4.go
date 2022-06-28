package main

//type money struct {
//	amount int
//	sync.Mutex
//}
//
//func (m *money) deposit(amount int) {
//	m.amount += amount
//}
//
//func main() {
//	mymoney := money{}
//	var wg sync.WaitGroup
//	wg.Add(1000)
//	i := 1
//	for ; i <= 1000; i++ {
//		go func(i int) {
//			mymoney.Lock()
//			defer mymoney.Unlock()
//			if i%2 == 0 {
//				mymoney.deposit(i)
//			} else {
//				mymoney.deposit(-i)
//			}
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//	fmt.Println("i: ", i)
//	fmt.Println(mymoney.amount)
//}
