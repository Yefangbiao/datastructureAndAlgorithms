package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("usage: add params filename1 filename2...")
		fmt.Println("example: ./main ./filename1 ./filename2")
	}
	for _, fileName := range files {
		f, err := os.Open(fileName)
		// 记得关闭文件
		var fc *os.File
		fc = f
		defer fc.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
		countLines(f)
	}
}

func countLines(f *os.File) {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// 有重复行就打印文件名字
	for _, v := range counts {
		if v > 1 {
			fmt.Println(f.Name())
			break
		}
	}
}
