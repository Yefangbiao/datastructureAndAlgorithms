package main

func main() {
	file1 := NewFile("a")
	file2 := NewFile("b")
	file3 := NewFile("c")

	folder1 := NewFolder("folder1")
	folder2 := NewFolder("folder2")
	folder3 := NewFolder("folder3")

	folder1.Add(folder2)
	folder1.Add(file1)
	folder2.Add(folder3)
	folder2.Add(file2)
	folder2.Add(file3)

	folder1.print("")
}
