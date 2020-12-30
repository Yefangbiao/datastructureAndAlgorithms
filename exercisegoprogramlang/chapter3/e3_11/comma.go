package e3_11

import (
	"bytes"
	"strings"
)

func Comma(s string) string {
	// 能够处理浮点数
	strs := strings.Split(s, `.`)
	// 忽略错误处理
	integers := strs[0]
	decimals := strs[1]
	if len(integers) <= 3 {
		return s
	}
	buf := bytes.Buffer{}
	buf.WriteString(integers[:len(integers)%3])
	for i := len(integers) % 3; i < len(integers); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(integers[i : i+3])
	}
	// 把小数写回去
	buf.WriteByte('.')
	buf.WriteString(decimals)
	return buf.String()
}
