package main

import "fmt"

type folder struct {
	folderName string
	nodes      []inode
}

func (f *folder) clone() inode {
	cloneFolder := &folder{
		folderName: f.folderName,
	}
	for _, i := range f.nodes {
		copy := i.clone()
		f.nodes = append(f.nodes, copy)
	}
	return cloneFolder
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation, f.folderName)
	for _, i := range f.nodes {
		i.print(indentation + indentation)
	}
}
