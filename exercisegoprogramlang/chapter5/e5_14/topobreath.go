package main

import "fmt"

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

// 假的广度优先拓扑排序，以后再改
func main() {
	seen := make(map[string]bool)
	orders := make([]string, 0, len(prereqs))
	var topoBreath func(list []string)
	topoBreath = func(list []string) {
		for _, l := range list {
			if !seen[l] {
				seen[l] = true
				topoBreath(prereqs[l])
				orders = append(orders, l)
			}
		}
	}

	keysMap := make(map[string]struct{})
	for key := range prereqs {
		keysMap[key] = struct{}{}
	}

	var keys []string
	for key := range keysMap {
		keys = append(keys, key)
	}

	topoBreath(keys)

	for i, course := range orders {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
