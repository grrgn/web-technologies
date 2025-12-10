// https://www.codewars.com/kata/5679aa472b8f57fb8c000047

package main

import "fmt"

func FindEvenIndex(arr []int) int {
	totalSum := 0
	for _, num := range arr {
		totalSum += num
	}
	leftSum := 0
	for i, currentNum := range arr {
		rightSum := totalSum - leftSum - currentNum
		if leftSum == rightSum {
			return i
		}
		leftSum += currentNum
	}
	return -1
}

func main() {
	arr1 := []int{1, 2, 3, 4, 3, 2, 1}
	fmt.Println(FindEvenIndex(arr1)) // 3

	arr2 := []int{20, 10, -80, 10, 10, 15, 35}
	fmt.Println(FindEvenIndex(arr2)) // 0

	arr3 := []int{1, 2, 3}
	fmt.Println(FindEvenIndex(arr3)) // -1
}