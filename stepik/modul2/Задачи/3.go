package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	var max rune
	split := []rune(str)
	for _, v := range split{
		if v > max{
			max = v
		}
	}
	fmt.Println(string(max))
}




