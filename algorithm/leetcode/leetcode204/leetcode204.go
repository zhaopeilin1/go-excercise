package leetcode204

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	} else {
		primes := []int{2}
		for i := 3; i < n; i++ {
			check := isPrime(primes, i)
			if check {
				primes = append(primes, i)
			}
		}
		return len(primes)
	}
}

func isPrime(primes []int, n int) bool {
	for _, prime := range primes {
		if n%prime == 0 {
			return false
		}
		if prime*prime > n {
			return true
		}
	}
	return true
}
