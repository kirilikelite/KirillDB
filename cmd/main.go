package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findPrimesInRange(start, end int, primes *[]int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; i <= end; i++ {
		if isPrime(i) {
			*primes = append(*primes, i)
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	start := 1
	end := 20000000

	now := time.Now()
	numGoroutines := runtime.NumCPU() * 1
	rangeSize := (end - start) / numGoroutines

	var wg sync.WaitGroup
	var primes []int

	for i := 0; i < numGoroutines; i++ {

		goroutineStart := start + i*rangeSize
		goroutineEnd := start + (i+1)*rangeSize - 1
		if i == numGoroutines-1 {
			goroutineEnd = end
		}
		wg.Add(1)
		go findPrimesInRange(goroutineStart, goroutineEnd, &primes, &wg)
	}

	wg.Wait()

	fmt.Println(time.Since(now))
}
