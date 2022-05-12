package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var scanstr string
	fmt.Scan(&scanstr)
	scanstr = strings.Replace(scanstr, " ", "", -1)
	strlen := utf8.RuneCountInString(scanstr)

	runeString := []rune(scanstr)
	overString := []rune(scanstr)

	for i := range runeString {
		overString[i] = runeString[strlen-1-i]
	}

	if string(runeString) == string(overString) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}

}
