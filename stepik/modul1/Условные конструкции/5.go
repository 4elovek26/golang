package main

import "fmt"

func main() {
	year := 0
	fmt.Scanln(&year)
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
