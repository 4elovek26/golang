package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var scanstr string
	fmt.Scan(&scanstr)

	lenstr := utf8.RuneCountInString(scanstr)

	for i:=0; i < lenstr; i++ {
		if strings.Count(scanstr, string(scanstr[i])) > 1 {
			scanstr = strings.Replace(scanstr, string(scanstr[i]), "", -1)
			i = 0
			lenstr = utf8.RuneCountInString(scanstr)
		}
	}
	fmt.Println(scanstr)


}