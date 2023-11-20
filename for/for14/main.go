package main

import (
	"fmt"
)

func main() {
	var n, i, sum float64
	fmt.Scan(&n)
	sum = 0
	for i = 1; i <= (2 * n - 1); i += 2{
		sum += i
	}
	fmt.Println(sum)
}