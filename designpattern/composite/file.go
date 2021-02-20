package main

import "fmt"

type file struct {
	fileName string
}

func (f *file) print(path string) {
	fmt.Printf("%s - %s\n", path, f.fileName)
}

func NewFile(fileName string) *file {
	return &file{fileName: fileName}
}
