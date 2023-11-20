package main

import (
	"fmt"
)

func main() {
	var n, i, sum float64
	fmt.Scan(&n)
	sum = 1
	add := 0.1
	for i = 1.1; i <= n; i += add {
		sum *= i
	}
	fmt.Println(sum)
}