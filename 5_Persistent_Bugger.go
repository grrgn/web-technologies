// https://www.codewars.com/kata/55bf01e5a717a0d57e0000ec

package main

import "fmt"

func Persistence(n int) int {
	steps := 0
	for n >= 10 {
		product := 1
		temp := n
		for temp > 0 {
			digit := temp % 10
			product = product * digit
			temp = temp / 10
		}
		n = product
		steps++
	}

	return steps
}

func main() {
	fmt.Println(Persistence(39)) // 3

	fmt.Println(Persistence(999)) // 4

	fmt.Println(Persistence(4)) // 0
}