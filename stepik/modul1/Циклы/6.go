package main

import "fmt"

func main() {
	n:= 0

	for fmt.Scan(&n);n <= 100 ;fmt.Scan(&n){
		if n >= 10 {
			fmt.Println(n)
		}
	}
}




