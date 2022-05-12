package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("text1.txt")
	reader := bufio.NewReaderSize(file, 65536)

	for i := 1; ; i++ {
		line, _ := reader.ReadString(';')
		if line == "0;" {
			fmt.Println(i)
		}
	}

}
