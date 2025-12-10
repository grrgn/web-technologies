// https://www.codewars.com/kata/550498447451fbbd7600041c

package main

import (
	"fmt"
	"sort"
)

func Comp(a []int, b []int) bool {
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		a[i] = a[i] * a[i]
	}
	sort.Ints(a)
	sort.Ints(b)

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func main() {
	a := []int{121, 144, 19, 161, 19, 144, 19, 11}
	b := []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}
	
	result := Comp(a, b)
	fmt.Println("Result:", result) // true
}