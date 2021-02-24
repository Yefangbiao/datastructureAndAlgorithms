package main

type tree struct {
	x, y int
	*treeType
}

func newTree(x int, y int, treeType *treeType) *tree {
	return &tree{x: x, y: y, treeType: treeType}
}

func (t *tree) draw() {
	t.treeType.draw(t.x, t.y)
}
