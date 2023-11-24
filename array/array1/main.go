package main

import "fmt"

func main() {
	var n, cnt int
	fmt.Scan(&n)
	a := make([]int, n+3)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		if ((a[i] > 0 && a[i + 1] > 0) || (a[i] < 0 && a[i + 1] < 0)){
			cnt++
		}
	}
	if cnt > 0{
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
	}
}