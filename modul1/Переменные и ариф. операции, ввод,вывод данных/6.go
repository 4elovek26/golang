package main

import "fmt"

func main(){

	var a int
	fmt.Scan(&a) // считаем переменную 'a' с консоли

	min := ((12*60)/360)*a

	fmt.Println("It is", min/60, "hours", (min%60), "minutes.")
}




