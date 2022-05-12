package main

import "fmt"


var m []int
func declOfNum(n int) int {
	j := 2
	m = append(m, 1)
	m = append(m, 1)

	for {

		m = append(m,m[j-1] + m[j-2])
		if m[j] >= n {
			break
		}
		j += 1
	}

	for i:= 0; i <= j; i++{
		if m[i] == n {
			return i+1
		}
	}
	return -1
}

func main(){
	var value int
	fmt.Scan(&value)
	val := declOfNum(value)
	fmt.Println(val)
}




