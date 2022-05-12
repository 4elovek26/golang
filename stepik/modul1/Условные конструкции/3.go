package main

import "fmt"

func main() {
	chislo := 0
	fmt.Scanln(&chislo)

	switch {
	case chislo == 10000:
		fmt.Println(1)
	case chislo >= 1000 && chislo < 10000:
		fmt.Println(chislo / 1000)
	case chislo >= 100 && chislo < 1000:
		fmt.Println(chislo / 100)
	case chislo >= 10 && chislo < 100:
		fmt.Println(chislo / 10)
	default:
		fmt.Println(chislo)
	}

}
