package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	ba := (a / 10 ) + (a % 10 )*10
	fmt.Println(ba)
}