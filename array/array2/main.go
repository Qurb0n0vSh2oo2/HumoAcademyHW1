package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+3)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	for i := n; i >= 1; i-- {
		fmt.Printf("%v ", a[i])
	}
	
}