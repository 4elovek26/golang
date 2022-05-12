package main

import "fmt"

func main() {
	n, c, count:= 0, 0, 1
	first := make(map[int]int)
	second := make(map[int]int)
	output := make(map[int]int)
	fmt.Scan(&n)
	fmt.Scan(&c)
	for i:=1; n >0; {
		first[i] = n%10
		n = n/10
		i++
	}
	for i:=1; c > 0; {
		second[i] = c%10
		c = c/10
		i++
	}

	for i := len(second); i>0; i--{
		for  j := len(first); j>0; j--{
			if first[j] == second[i]{
				output[count] = first[j]
				count++
			}
		}
	}
	for i := len(first); i>0; i--{
		for  j := len(output); j>0; j--{
			if first[i] == output[j]{
				fmt.Print(output[j]," ")
			}
		}
	}
}




