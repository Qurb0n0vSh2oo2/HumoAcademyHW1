package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var res bool
	res = (a > 2) && (b <= 3)
	fmt.Println(res)
}