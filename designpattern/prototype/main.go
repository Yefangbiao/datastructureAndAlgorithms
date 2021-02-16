package main

func main() {
	file1 := &file{
		filename: "file1",
	}
	file2 := &file{
		filename: "file2",
	}
	f1 := &folder{
		folderName: "f1",
		nodes:      []inode{file1},
	}
	f2 := &folder{
		folderName: "f2",
		nodes:      []inode{f1, file2},
	}
	f2.print("-")
}
