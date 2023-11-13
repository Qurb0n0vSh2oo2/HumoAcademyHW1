package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if (a > 0) {
		fmt.Println(a + 1)
	} else if ( a < 0 ) {
		fmt.Println(a - 2)
	} else {
		fmt.Println(10)
	}
}