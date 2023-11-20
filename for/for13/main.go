package main

import (
	"fmt"
)

func main() {
	var n, i, sum1, sum2 float64
	fmt.Scan(&n)
	sum1 = 0
	sum2 = 0
	oddeven := 0.2
	for i = 1.1; i <= n; i += oddeven {
		sum1 += i
	}
	for i = 1.2; i <= n; i += oddeven {
		sum2 += i

	}
	fmt.Println(sum1 + (sum2 * (-1)))
}