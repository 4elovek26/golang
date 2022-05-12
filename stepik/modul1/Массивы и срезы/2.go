package main

import "fmt"
var value int
func main() {
	count := 0
	fmt.Scanln(&count)
	slice := []int{}
	for fmt.Scan(&value);count >0; fmt.Scan(&value){
		slice = append(slice, value)
		count--
	}
	fmt.Println(slice[3])
}