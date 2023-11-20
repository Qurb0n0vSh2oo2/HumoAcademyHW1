package main

import (
	"fmt"
	"math"
)

func main() {
	var n, i, sum float64
	fmt.Scan(&n)
	sum = 0
	for i = 0; i <= n; i++ {
		sum += math.Pow((n + i), 2)
	}
	fmt.Println(sum)
}