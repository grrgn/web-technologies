// https://www.codewars.com/kata/5626b561280a42ecc50000d1

package main

import (
	"fmt"
	"math"
	"strconv"
)

func SumDigPow(a, b int) []int {
	var result []int
	for i := a; i <= b; i++ {
		if isEureka(i) {
			result = append(result, i)
		}
	}

	return result
}

func isEureka(n int) bool {
	s := strconv.Itoa(n)
	
	var sum float64
	for i, char := range s {
		digit := float64(char - '0')
		power := float64(i + 1)
		sum += math.Pow(digit, power)
	}
	return int(sum) == n
}

func main() {
	fmt.Println(SumDigPow(1, 10)) 
	// [1 2 3 4 5 6 7 8 9]

	fmt.Println(SumDigPow(1, 100)) 
	// [1 2 3 4 5 6 7 8 9 89]

	fmt.Println(SumDigPow(10, 150)) 
	// [89 135]
}