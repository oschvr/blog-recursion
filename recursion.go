package main

import (
	"fmt"
)

func main() {
	fmt.Println(recursion(10))
}

func recursion(x int) int {
	if x == 0 {
		return 0
	}
	fmt.Print(x, ", ")
	return recursion(x - 1)
}
