package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var f bool
	a := make([]int, n+2)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	// cnt := 0
	min := 0
	max := 1001
	for i := 0; i < n; i++ {
		if a[i] > min {
			max = a[i]
		}
		if a[i] < max || !f {
			min = a[i]
			f = true
		}

	}
	for i := 0; i < n; i++ {
		if a[i] == max {
			a[i] = min
		}
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%v ", a[i])
	}
	// fmt.Println(max, min)

}
