package main

import "fmt"

var a1, a2, a3 int

func main() {
	chislo := 0

	fmt.Scanln(&chislo)
	a1 = chislo % 10
	a2 = chislo % 100 / 10
	a3 = chislo / 100
	if a1 != a2 && a1 != a3 && a2 != a3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
