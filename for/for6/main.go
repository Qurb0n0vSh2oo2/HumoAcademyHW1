package main

import "fmt"

func main() {
	var n, i float64
	fmt.Scan(&n)
	for i = 1.2; i <= 2; i+=0.2 {
		fmt.Println(n * i)
	}
}