package chapter16

import "testing"

func TestHuffman(t *testing.T) {
	char := newChar()
	char.Insert('f', 5)
	char.Insert('e', 9)
	char.Insert('c', 12)
	char.Insert('b', 13)
	char.Insert('d', 16)
	char.Insert('a', 45)

	root := Huffman(*char)
	printTree(root, "")
}
