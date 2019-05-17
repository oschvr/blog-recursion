package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// fmt.Println(iterativeFactorial(14))
	fmt.Println(factorial(10))
	fmt.Println(time.Since(start))
}

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func iterativeFactorial(n int64) int64 {
	var result, i int64 = 1, 1
	for ; i <= n; i++ {
		result = result * i
	}
	return result
}
