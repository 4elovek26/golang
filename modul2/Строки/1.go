package main

import (
	"fmt"
	"unicode"
	"bufio"
	"os"
	"unicode/utf8"
)

var strlen int

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	mytext := []rune(text)
	strlen = utf8.RuneCountInString(text)
	if unicode.IsUpper(mytext[0]) && mytext[strlen-1] == '.'{
		fmt.Println("Right")
	}else{
		fmt.Println("Wrong")
	}
}




