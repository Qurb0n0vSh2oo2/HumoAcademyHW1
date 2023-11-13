package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var a1, a3 int
	a1 = a / 100
	a3 = a % 10
	fmt.Println(a1, " " , a3)
	if a1 > a3 {
		fmt.Println(a1)
	} else {
		fmt.Println(a3)
	}

}