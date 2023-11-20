package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	n := 0
	for i := b-1; i > a; i-- {
		fmt.Println(i)
		n++
	}
	fmt.Println("Kolichestvo", n)
}