package main

import "fmt"

func main() {
	var n, i float64
	fmt.Scan(&n)
	for i = 1; i <= 10; i++ {
		fmt.Println(i * n)
	}
}