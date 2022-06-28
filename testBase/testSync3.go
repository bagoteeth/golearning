package main

//var money int
//
//func deposit(amount int) {
//	money += amount
//}
//
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(1000)
//	i := 1
//	ch := make(chan struct{}, 1)
//	for ; i <= 1000; i++ {
//		go func(i int) {
//			ch <- struct{}{}
//			if i%2 == 0 {
//				deposit(i)
//			} else {
//				deposit(-i)
//			}
//			wg.Done()
//			<-ch
//		}(i)
//	}
//	wg.Wait()
//	fmt.Println("i: ", i)
//	fmt.Println(money)
//}
