package main

import "fmt"

func main() {
	chislo := 0
	fmt.Scanln(&chislo)
	switch {
	case chislo > 0: fmt.Println("Число положительное")
	case chislo < 0: fmt.Println("Число отрицательное")
	default: fmt.Println("Ноль")
	}
}