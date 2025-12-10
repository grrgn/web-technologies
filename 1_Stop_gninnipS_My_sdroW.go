// https://www.codewars.com/kata/5264d2b162488dc400000001

package main

import (
	"fmt"
	"strings"
)

func SpinWords(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		if len(word) >= 5 {
			words[i] = reverseString(word)
		}
	}
	return strings.Join(words, " ")
}

func reverseString(s string) string {
	runes := []rune(s)
	
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	
	return string(runes)
}

func main() {
	fmt.Println(SpinWords("Hey fellow warriors"))  // Hey wollef sroirraw
	fmt.Println(SpinWords("This is a test"))       // This is a test
	fmt.Println(SpinWords("This is another test")) // This is rehtona test
}