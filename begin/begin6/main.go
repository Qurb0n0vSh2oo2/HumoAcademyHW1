package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Print(a * b * c, " ", 2 * (a * b + b * c + a * c))
}