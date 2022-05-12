package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	s = strings.Replace(s, " ", "", -1)
	index := strings.Index(s, ";")
	runeWord := []rune(s)
	var left, right string
	for i, v := range runeWord {
		if i < index {
			left += string(v)
		}
		if i > index {
			right += string(v)
		}
	}
	left = strings.Replace(left, ",", ".", -1)
	right = strings.Replace(right, ",", ".", -1)
	resultLeft, err := strconv.ParseFloat(left, 64)
	if err != nil {
		panic(err)
	}
	resultRight, err := strconv.ParseFloat(right, 64)
	if err != nil {
		panic(err)
	}
	result := resultLeft / resultRight
	fmt.Printf("%.4f", result)

}
