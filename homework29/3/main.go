package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < len(a); i+= m {
		b := i + m
		if b > len(a) {
			b = len(a) 
		}
		fmt.Println(a[i:b])
	}

	
}
