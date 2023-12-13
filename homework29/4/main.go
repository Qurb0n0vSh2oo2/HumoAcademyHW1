package main

import (
	"fmt"
)

func main() {
	var n int
	// var f bool
	fmt.Scan(&n)
	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}
	max := 1
	for _, value := range ar {
		max = value * value
		// fmt.Printf("%v ", max)
	}
	
	// fmt.Println(max, min)
	ar2 := make([]int, 0)
	for _, value := range ar {
		if value == max {
			ar2 = append(ar2, value)
		}
		break

	}
	fmt.Println(ar2)
}

// Удаление элемента из массива

// Удаление всех положительных чётных, и отрицательных нечётных элементов
// n

// [-1, 2, -3, -4, 5]
//   0, 1,  2,  3, 4
// i = 0, [2, -3, -4, 5]
//         0,  1,  2, 3
// i = 1
