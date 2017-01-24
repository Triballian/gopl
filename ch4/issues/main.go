package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Triballian/gopl/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	// exmaple test run: go run main.go repo:golang/go is:open json decoder
}
