package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	split := strings.Split(str, "")
	fmt.Println(strings.Join(split, "*"),)
}





