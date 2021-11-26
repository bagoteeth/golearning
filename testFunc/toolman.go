package testFunc

func TryCover() int {
	sum := 0
	for i := 10; i < 50; i++ {
		if i < 0 {
			sum -= i
		} else if i == 0 {
			sum *= 10
		} else if i > 0 {
			sum += i
		}
	}
	return sum
}

func TryBenchmark() int {
	sum := 0
	for i := 0; i <= 1000; i++ {
		sum += i
	}
	return sum
}

func TryFastBenchmark() int {
	return 500500
}
