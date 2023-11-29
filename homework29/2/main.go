package main

import "fmt"

func main() {
	var n int
	//   var f bool
	fmt.Scan(&n)
	ar := make([]int, n)
	// a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	// a[0] = ar[0]
	// a[1] = ar[n - 1]
	// a[2] = ar[n - 4]
	// a[3] = ar[n - 2]
	// a[4] = ar[n - 3]
	// fmt.Println(a)

	left := 0
	right := n - 1
	for i := 0; i < n; i++ {
		if left <= right {
			fmt.Printf("%v ", ar[left])

			if left != right {
				fmt.Printf("%v ", ar[right])
			}
			left++
			right--
		}
	}

}
