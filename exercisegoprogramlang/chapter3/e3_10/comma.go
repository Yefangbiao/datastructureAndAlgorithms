package e3_10

import "bytes"

// 编写一个非递归的comma函数，运用bytes.Buffer

func Comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	buf := bytes.Buffer{}
	buf.WriteString(s[:len(s)%3])
	for i := len(s) % 3; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
