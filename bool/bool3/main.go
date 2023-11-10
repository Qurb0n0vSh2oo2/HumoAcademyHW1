package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var b bool
	b = a % 2 == 0 
	fmt.Println(b)

}