package main

import (
	"fmt"
)

func main() {
	var n, min int
	var f bool
	fmt.Scan(&n)
	a := make([]int, n + 2)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if (a[i] % 2 != 0 && (a[i] < min || !f) ) {
			min = a[i]
			f = true
		}
	}

	if f {
		fmt.Println(min)
	} else{
		fmt.Println(0)
	}
}
	
// var n int
	// // var m int
	// fmt.Scan(&n)
	// // k := 1
	// a := make([]int, n+3)
	// b := make([]int, n + 1)
	// // c := make([]int, n)
	
	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&a[i])
	// }
	// for i := 0; i < n; i++ {
	// 	if ( a[i] % 2 != 0 ){
	// 		b[i] = a[i]
	// 	}
	// }
	// sort.Ints(b)
	// for i := 0; i < n; i++ {

	// 		if ( b[i] < b[n] && b[i] != 0 ){
	// 			fmt.Printf("%v ", b[i])
	// 		}else{
	// 			fmt.Println(0)
	// 			break
	// 		}
