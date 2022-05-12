package main

import (
	"fmt"
	"strings"
)

func main() {
	var scanstr string
	fmt.Scan(&scanstr)

	strSplit := strings.Split(scanstr, "")

	for i := range strSplit {
		if i%2 != 0 {
			fmt.Print(strSplit[i])
		}
	}


}




