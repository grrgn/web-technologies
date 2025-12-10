// https://www.codewars.com/kata/57eb8fcdf670e99d9b000272

package main

import (
	"fmt"
	"strings"
)

func High(s string) string {
	words := strings.Split(s, " ")

	highestScore := 0
	winnerWord := ""

	for _, word := range words {
		currentScore := 0
		for _, char := range word {
			score := int(char - 'a' + 1)
			currentScore += score
		}
		if currentScore > highestScore {
			highestScore = currentScore
			winnerWord = word
		}
	}

	return winnerWord
}

func main() {
	fmt.Println(High("man i need a taxi up to ubud")) 
	// taxi

	fmt.Println(High("what time are we climbing up the volcano")) 
	// what

	fmt.Println(High("aa b")) 
	// aa
}