package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	timeStr, err := rd.ReadString('\n')
	if err != nil {
		//...
	}

	// Удаление последнего символа при помощи пакета "strings".
	timeStr = strings.TrimSuffix(timeStr, "\n")

	firstTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		panic(err)
	}


	if firstTime.Hour() > 13{
		firstTime = firstTime.AddDate(0,0,1)
		fmt.Println(firstTime.Format("2006-01-02 15:04:05"))
	}else{
		fmt.Println(firstTime.Format("2006-01-02 15:04:05"))
	}



}