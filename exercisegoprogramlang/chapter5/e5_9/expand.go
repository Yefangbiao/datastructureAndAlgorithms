package e5_9

import "strings"

func Expand(s string, f func(string) string) string {
	ret := strings.Replace(s, "$foo", f("foo"), -1)
	return ret
}

func replace(string) string {
	//将foo全部替换为"abc"
	return "abc"
}
