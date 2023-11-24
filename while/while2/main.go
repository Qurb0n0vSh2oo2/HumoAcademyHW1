package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	l := a
	cnt := 0
	for l >= b {
		l -= b
		cnt++
	}
	
	fmt.Println(cnt)
}