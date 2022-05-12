package main

import "fmt"

func main() {
	count, newInt, sum := 0, 0, 0
	fmt.Scanln(&count)
	for fmt.Scan(&newInt);count > 0;fmt.Scan(&newInt) {
		if newInt < 100 && newInt % 8 == 0 && newInt >= 10 {
			sum += newInt
		}
		count--
	}
	fmt.Println(sum)
}




