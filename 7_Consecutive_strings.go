// https://www.codewars.com/kata/56a5d994ac971f1ac500003e

package main

import (
	"fmt"
	"strings"
)

func LongestConsec(strarr []string, k int) string {
	n := len(strarr)
	if n == 0 || k > n || k <= 0 {
		return ""
	}

	longest := ""
	for i := 0; i <= n-k; i++ {
		currentStr := strings.Join(strarr[i : i+k], "")
		if len(currentStr) > len(longest) {
			longest = currentStr
		}
	}

	return longest
}

func main() {
	arr1 := []string{"tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"}
	fmt.Println(LongestConsec(arr1, 2)) 
	// folingtrashy

	arr2 := []string{"zone", "abigail", "theta", "form", "libe", "zas"}
	fmt.Println(LongestConsec(arr2, 2)) 
	// abigailtheta

	fmt.Println(LongestConsec(arr1, 100)) 
	// ""
}