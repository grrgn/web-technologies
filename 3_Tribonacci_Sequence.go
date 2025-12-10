// https://www.codewars.com/kata/556deca17c58da83c00002db

package main

import "fmt"

func Tribonacci(signature [3]float64, n int) []float64 {
	result := signature[:]

	if n == 0 {
		return []float64{}
	}
	if n <= 3 {
		return result[:n]
	}

	for i := 3; i < n; i++ {
		nextNum := result[i-1] + result[i-2] + result[i-3]
		result = append(result, nextNum)
	}

	return result
}

func main() {
	sig1 := [3]float64{1, 1, 1}
	fmt.Println(Tribonacci(sig1, 10))
	sig2 := [3]float64{0, 0, 1}
	fmt.Println(Tribonacci(sig2, 10))
	fmt.Println(Tribonacci(sig1, 1)) 
}