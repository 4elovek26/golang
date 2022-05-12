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
	for ind := range slice{
		if ind % 2 == 0 {
			fmt.Print(slice[ind]," ")
		}
	}
}




