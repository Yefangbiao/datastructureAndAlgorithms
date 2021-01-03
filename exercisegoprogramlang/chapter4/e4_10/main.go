package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	// its[0]一个月内，its[1]一年内，its[2]超过一年
	its := make([][]*Issue, 3)
	for i := range its {
		its[i] = []*Issue{}
	}
	// 一个月前的时间
	monthAgo := time.Now().AddDate(0, -1, 0)
	// 一年前的时间
	yearAgo := time.Now().AddDate(-1, 0, 0)
	for _, item := range result.Items {
		if monthAgo.Before(item.CreatedAt) {
			// 一个月内
			its[0] = append(its[0], item)
		} else if yearAgo.Before(item.CreatedAt) {
			its[1] = append(its[1], item)
		} else {
			its[2] = append(its[2], item)
		}
	}

	strs := []string{"一个月内", "一年内", "一年前"}

	for i, it := range its {
		fmt.Println("---------", strs[i], "---------")
		for _, item := range it {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
