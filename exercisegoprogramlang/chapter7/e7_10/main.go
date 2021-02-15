package main

import (
	"fmt"
	"sort"
)

type demo struct {
	nums []int
}

func (d *demo) Len() int {
	return len(d.nums)
}

func (d *demo) Less(i, j int) bool {
	return d.nums[i] < d.nums[j]
}

func (d *demo) Swap(i, j int) {
	d.nums[i], d.nums[j] = d.nums[j], d.nums[i]
}

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
	}
	return true
}

func main() {
	d := demo{[]int{1, 2, 1}}
	fmt.Println(IsPalindrome(&d))
	fmt.Println(IsPalindrome(&demo{[]int{1, 2, 3}}))
}
