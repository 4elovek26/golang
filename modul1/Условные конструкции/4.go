package main

import "fmt"



func main() {
	chislo := 0
	fmt.Scanln(&chislo)
	left, right := chislo / 1000, chislo % 1000
	luckyLeft, luckyRight := 0, 0
	for left != 0 && right != 00{
		luckyLeft += left % 10
		luckyRight += right % 10
		left = left / 10
		right = right / 10
	}
	if luckyLeft == luckyRight {
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
	}
}




