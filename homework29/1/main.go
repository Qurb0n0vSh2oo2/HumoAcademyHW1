package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	// var f bool
	fmt.Scan(&n)
	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}
	min := math.MaxInt
	max := math.MinInt
	
	for i := 0; i < n; i++ {
		if ar[i] > 0 && ar[i] < min{
			min = ar[i]
		}
		if ar[i] < 0 && ar[i] > max {
			max = ar[i]
		}
	}
	// fmt.Println(max, min)

	for i, value := range ar {
		if value == min {
			ar = append(ar[:i], ar[i+1:]...)
		}
		break
		
	}
	for i, value := range ar {
		if value == max {
			ar = append(ar[:i], ar[i+1:]...)
		}
		break
	}
	fmt.Println(ar)
}

// Удаление элемента из массива

// Удаление всех положительных чётных, и отрицательных нечётных элементов
// n

// [-1, 2, -3, -4, 5]
//   0, 1,  2,  3, 4
// i = 0, [2, -3, -4, 5]
//         0,  1,  2, 3
// i = 1

