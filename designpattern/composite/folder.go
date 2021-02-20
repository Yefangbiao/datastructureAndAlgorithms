package main

import "fmt"

type folder struct {
	folderName string
	com        []component
}

func (f *folder) Add(c component) {
	f.com = append(f.com, c)
}

func (f *folder) print(path string) {
	fmt.Printf("%s - %s\n", path, f.folderName)
	for _, c := range f.com {
		c.print(path + " - ")
	}
}

func NewFolder(folderName string) *folder {
	return &folder{folderName: folderName}
}
