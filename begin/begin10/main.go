package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println("Summa kvadratov", a*a + b*b)
	fmt.Println("Raznost kvadratov", a*a - b*b)
	fmt.Println("Proizvedenie kvadratov", a*a * b*b)
	fmt.Println("Chastnoe kvadratov", a*a - b*b)
	
}