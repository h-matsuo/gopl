package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/h-matsuo/gopl/ch04/ex10/github"
)

func main() {
	page := 1
	result, err := github.SearchIssues(os.Args[1:], page)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for {
		for _, item := range result.Items {
			printPeriodIfNeeded(item.CreatedAt)
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
		// Consider pagenation
		if result.IncompleteResults {
			page++
			result, _ = github.SearchIssues(os.Args[1:], page)
			continue
		}
		break
	}
}

var (
	now            = time.Now()
	oneYearBefore  = now.AddDate(-1, 0, 0)
	oneMonthBefore = now.AddDate(0, -1, 0)

	printedMoreThanOneYearAgo  = false
	printedLessThanOneYearAgo  = false
	printedLessThanOneMonthAgo = false
)

func printPeriodIfNeeded(t time.Time) {
	if !printedMoreThanOneYearAgo && t.Before(oneYearBefore) {
		fmt.Println("\n===== [Created more then a year ago] =====")
		printedMoreThanOneYearAgo = true
	} else if !printedLessThanOneYearAgo && t.Before(oneMonthBefore) {
		fmt.Println("\n===== [Created within a year] =====")
		printedLessThanOneYearAgo = true
	} else if !printedLessThanOneMonthAgo {
		fmt.Println("\n===== [Created within a month] =====")
		printedLessThanOneMonthAgo = true
	}
}
