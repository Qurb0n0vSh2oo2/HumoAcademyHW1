package main

import (
	"fmt"
)

func main() {
	var n, a, i, sum float64
	fmt.Scan(&a, &n)
	sum = 1
	for i = 1; i <= n; i++ {
		sum *= a
	}
	fmt.Println(sum)
}