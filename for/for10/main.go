package main

import "fmt"

func main() {
	var n, i, sum float64
	fmt.Scan(&n)
	sum = 0
	for i = 2; i <= n; i++ {
		sum += (1/i)
	}
	fmt.Println(1+sum)
}