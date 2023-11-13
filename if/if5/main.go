package main

import "fmt"

func main() {
	var a, b, c, res1, res2 int
	res1 = 0
	res2 = 0
	fmt.Scan(&a, &b, &c)
	if ( a > 0 ){
		res1 += 1
	} else {
		res2 += 1
	}
	if ( b > 0 ){
		res1 += 1
	}else {
		res2 += 1
	}
	if ( c > 0 ){
		res1 += 1
	}else {
		res2 += 1
	}
	fmt.Println(res1, " ", res2)
}