package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var m int
	fmt.Scan(&n)
	// k := 1
	a := make([]int, n+3)
	b := make([]int, n)
	// c := make([]int, n)
	
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		if ( a[i] % 2 != 0 ){
			b[i] = a[i]
		}
	}
	sort.Ints(b)
	for i := 0; i < n; i++ {

			if ( b[i] > 0 ){
				m = b[i]
				break
			}
			break

	}
	
	fmt.Println(b)
	
	// fmt.Println(c)
	
	fmt.Println(m)
	
}