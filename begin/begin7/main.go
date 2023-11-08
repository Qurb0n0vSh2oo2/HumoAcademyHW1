package main

import "fmt"

func main() {
	var r float64
	var pi float64 = 3.14
	fmt.Scan(&r)
	fmt.Print(2 * pi * r, " ", pi * r * r)
}