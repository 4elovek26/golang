package main

import "fmt"

func main() {
	newInt, sum, big:= 0, 0, 0
	for fmt.Scan(&newInt);newInt > 0;fmt.Scan(&newInt) {
		if big < newInt {
			big = newInt
			sum = 0
		}
		if big == newInt {
			sum += 1
		}
	}
	fmt.Println(sum)
}




