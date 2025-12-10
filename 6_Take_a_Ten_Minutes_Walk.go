// https://www.codewars.com/kata/54da539698b8a2ad76000228

package main

import "fmt"

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}
	x := 0
	y := 0
	for _, direction := range walk {
		switch direction {
		case 'n':
			y++
		case 's':
			y--
		case 'e':
			x++
		case 'w':
			x--
		}
	}

	if x == 0 && y == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	walk1 := []rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}
	fmt.Println(IsValidWalk(walk1)) // true

	walk2 := []rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'w', 'w', 'w'}
	fmt.Println(IsValidWalk(walk2)) // false

	walk3 := []rune{'n', 'n', 'n', 's', 'n', 's', 'n', 's', 'n', 's'} 
	fmt.Println(IsValidWalk(walk3)) // false
}