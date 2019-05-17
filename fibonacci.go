package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	n := 10
	fmt.Println(fibonacci(n))

	fmt.Println(time.Since(start))

}

// Recursive fibonacci
// function definition
func fibonacci(n int) int {
	// Base case
	if n <= 1 {
		return n
	}
	// F(n) = F(n-1) + F(n-2)
	return fibonacci(n-1) + fibonacci(n-2)
}

// Iterative fibonacci
// using slices
func fibonacciLoop(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
