package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	/// 3 2
	if (a > b) {
		a = a + b; ///a =3+2 = 5
		b = a - b; ///b =5-2 = 3
		a = a - b; ///a =5-3 = 2
		fmt.Println(a, " ", b)
	} else {
		fmt.Println(a, " ", b)
	}
}