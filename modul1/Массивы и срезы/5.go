package main

import "fmt"



func main() {
	count := 0
	fmt.Scanln(&count)
	slice := []int{}
	for fmt.Scan(&value);count >0; fmt.Scan(&value){
		slice = append(slice, value)
		count--
	}
	sum := 0
	for _, val := range slice{
		if val >= 0 {
			sum += 1
		}
	}
	fmt.Print(sum)
}