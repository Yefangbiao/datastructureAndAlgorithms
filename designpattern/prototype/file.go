package main

import "fmt"

type file struct {
	filename string
}

func (f *file) clone() inode {
	clone := &file{
		filename: f.filename,
	}
	return clone
}

func (f *file) print(indentation string) {

	fmt.Println(indentation, f.filename)
}
