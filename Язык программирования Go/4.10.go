package main

import (
	"fmt"
	"gopl.io/ch4/github"
	"log"
	"os"
	"time"
)

//Program arguments: ./4.10 repo:golang/go is:open json decoder
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("Поданы около года назад")
	for _, item := range result.Items {
		if (time.Now().Year() - item.CreatedAt.Year()) <= 1 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Println("Поданы более одного года назад")
	for _, item := range result.Items {
		if (time.Now().Year() - item.CreatedAt.Year()) > 1 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
