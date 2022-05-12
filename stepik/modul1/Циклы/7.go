package main

import "fmt"

func main() {
	n, c, d, years:= 0, 0, 0, 0
	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)

	for n <= d{
		n = n + (n*c)/100
		years += 1
	}
	fmt.Println(years)
}




