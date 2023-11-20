package main

import "fmt"

func main() {
	var n, i float64
	fmt.Scan(&n)
	for i = 0.1; i <= 1; i+=0.1 {
		fmt.Println(n * i)
	}
}