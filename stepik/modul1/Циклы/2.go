package main

import "fmt"

func main() {
	a, b,sum := 0,0,0
	fmt.Scanln(&a, &b)
	for ;a <= b; a++ {
		sum += a
	}
	fmt.Println(sum)
}




