package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a * a * a, " ", 6 * a * a)
}