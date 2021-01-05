package e5_16

import "strings"

func Joins(sep string, strs ...string) string {
	res := strings.Builder{}
	for i := 0; i < len(strs); i++ {
		if i > 0 {
			res.WriteString(sep)
		}
		res.WriteString(strs[i])
	}
	return res.String()
}
