package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var res bool
	res = (a >= 0) || (b < -2)
	fmt.Println(res)
}