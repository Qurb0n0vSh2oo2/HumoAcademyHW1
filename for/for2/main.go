package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	n := 0
	for i := a; i <= b; i++ {
		fmt.Println(i)
		n++
	}
	fmt.Println("Kolichestvo", n)
}