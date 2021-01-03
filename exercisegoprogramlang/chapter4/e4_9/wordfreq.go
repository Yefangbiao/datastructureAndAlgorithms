package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 在本地命令行执行
	f, err := os.Open("./test.txt")
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(bufio.NewReader(f))
	scanner.Split(bufio.ScanWords)

	words := make(map[string]int)
	for scanner.Scan() {
		word := scanner.Text()
		words[word]++
	}

	for k, v := range words {
		fmt.Println(k, v)
	}
}
