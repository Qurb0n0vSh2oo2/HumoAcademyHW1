package main

import "fmt"

func main() {
	var a, b, c, res int
	res = 0
	fmt.Scan(&a, &b, &c)
	if ( a > 0 ){
		res += 1
	}
	if ( b > 0 ){
		res += 1
	}
	if ( c > 0 ){
		res += 1
	}
	fmt.Println(res)
}