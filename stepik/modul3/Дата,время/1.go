package main

import (
	"fmt"
	"time"
)

func main() {
	var taim string
	fmt.Scan(&taim)

	firstTime, err := time.Parse(time.RFC3339, taim)
	if err != nil {
		panic(err)
	}

	//unixTime := timeRFC.Format()
	fmt.Println(firstTime.Format(time.UnixDate))

}