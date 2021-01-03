package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// 统计字母、数字和其他在unicode分类的字符数目
func main() {
	counts := make(map[string]int) // 各种分类的字符数目
	invalid := 0                   // 非法UTF-8字符数量
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() //rune、nbytes、error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stdout, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			counts["letter"]++
		} else if unicode.IsDigit(r) {
			counts["digit"]++
		} else {
			counts["others"]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
