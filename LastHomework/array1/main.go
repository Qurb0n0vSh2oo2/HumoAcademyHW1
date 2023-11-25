package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scan(&n)
	a := make([]int, n + 2)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&x)
	cnt := 0
	for i := 0; i < n; i++ {
		if x == a[i]{
			cnt++
		}
	}
	fmt.Println(cnt)

}
	