package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	split := []rune(str)
	for _, v := range split{
		tmp := v - 48
		fmt.Print(tmp*tmp)
	}
}




