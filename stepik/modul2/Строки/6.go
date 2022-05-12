package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	// put your code here
	var scanstr string
	fmt.Scan(&scanstr)
	yes := 1
	//runeString := []rune(scanstr)
	if utf8.RuneCountInString(scanstr) >= 5 {
		for _, r := range scanstr {
			if unicode.IsLetter(r) {
				if unicode.Is(unicode.Latin, r) {
					continue
				}
			}
			if unicode.IsDigit(r) {
				continue
			}
			fmt.Println("Wrong password")
			yes = 0
			break
		}
		if yes == 1 {
			fmt.Println("Ok")
		}
	}else{
		fmt.Println("Wrong password")
	}
}