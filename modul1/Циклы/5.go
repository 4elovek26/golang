package main

import "fmt"

func main() {
	n, c, d:= 0, 0, 0
	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)
	for i := 1; i <= n; i++{
		if i % c == 0 && i % d == 0 {
			continue
		}else if i % c == 0{
			fmt.Println(i)
			break
		}
	}
}




