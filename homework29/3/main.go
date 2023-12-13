package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
		if (i+1)%m == 0 {
			fmt.Print("\n")
		}
	}

}
