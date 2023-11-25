package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scan(&n)
	var f bool
	a := make([]int, n+2)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&x)
	cnt := 0
	min := 0
	max := 1001
	for i := 0; i < n; i++ {
		if x == a[i] {
			cnt = a[i]

		}
		if x != a[i] {
			if a[i] > min {
				max = a[i]
			}
			if a[i] < max || !f {
				min = a[i]
				f = true
			}
		}

	}
	// fmt.Println(cnt, max, min)
	if cnt > 0 {
		fmt.Println(cnt)
	} else {
		if x-min > 0 {
			fmt.Println(max)
		} else if x-max < 0 {
			fmt.Println(min)
		}
	}

}
