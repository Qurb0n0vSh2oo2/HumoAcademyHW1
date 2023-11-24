package main

import "fmt"

func main() {
	var n, cnt int
	fmt.Scan(&n)
	a := make([]int, n+3)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i < n - 1; i++ {
		if ( a[i] > a[i - 1] && a[i] > a[i + 1]){
			cnt++
		}
	}
	fmt.Println(cnt)
	
}