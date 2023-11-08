package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	sum := (a / 10 ) + (a % 10 )
	pro := (a / 10 ) * (a % 10 )
	fmt.Println(sum)
	fmt.Println(pro)
}