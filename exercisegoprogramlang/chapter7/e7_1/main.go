package main

import (
	"bufio"
	"fmt"
	"strings"
)

// 单词的计数器
type WordCounter int

func (w *WordCounter) Write(p []byte) (n int, err error) {

	scan := bufio.NewScanner(strings.NewReader(string(p)))
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		*w += 1
	}
	return len(p), nil
}

// 行计数器
type LineCounter int

func (w *LineCounter) Write(p []byte) (n int, err error) {

	scan := bufio.NewScanner(strings.NewReader(string(p)))
	scan.Split(bufio.ScanLines)

	for scan.Scan() {
		*w += 1
	}
	return len(p), nil
}

func main() {

	var c WordCounter
	fmt.Fprintf(&c, "hello, %s", "world")
	fmt.Println(c)

	var l LineCounter
	fmt.Fprintln(&l, "hello, %s", "world")
	fmt.Println(l)
}
