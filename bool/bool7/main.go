package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	var res bool
	res = ((a < b) && (b < c) ||(a > b) && (b > c))
	fmt.Println(res)
}