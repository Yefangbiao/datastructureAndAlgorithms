package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	orders := topoSort(prereqs)
	if !isLegal(orders) {
		fmt.Println("不合法")
		return
	}
	for i, course := range orders {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	keysMap := make(map[string]struct{})
	for key := range m {
		keysMap[key] = struct{}{}
	}

	var keys []string
	for key := range keysMap {
		keys = append(keys, key)
	}

	visitAll(keys)
	return order
}

func isLegal(order []string) bool {
	// 验证是否合法
	seen := make(map[string]bool)
	for _, course := range topoSort(prereqs) {
		seen[course] = true
		for _, pre := range prereqs[course] {
			if !seen[pre] {
				return false
			}
		}
	}
	return true
}
