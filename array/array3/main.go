package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	a := make([]int, n+3)
	m = 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		if a[i] > m{
			m = a[i]
		}
	}
	fmt.Println(m)
	
}