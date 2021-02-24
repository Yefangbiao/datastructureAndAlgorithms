package main

type forest struct {
	trees []*tree
}

func NewForest() *forest {
	return &forest{
		trees: []*tree{},
	}
}

func (f *forest) PlantTree(x, y int, name, color, texture string) {
	newType := SingleTreeFactory.getTreeType(name, color, texture)
	newTree := newTree(x, y, newType)
	f.trees = append(f.trees, newTree)
}

func (f *forest) Draw() {
	for _, tree := range f.trees {
		tree.draw()
	}
}
