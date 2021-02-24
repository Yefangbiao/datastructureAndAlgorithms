package main

type treeFactory struct {
	treeTypes []*treeType
}

var SingleTreeFactory *treeFactory

func init() {
	GetSingleTreeFactory()
}

func GetSingleTreeFactory() {
	if SingleTreeFactory == nil {
		SingleTreeFactory = &treeFactory{treeTypes: []*treeType{}}
	}
}

func (t *treeFactory) getTreeType(name, color, texture string) *treeType {
	tree := t.find(name, color, texture)
	if tree == nil {
		tree = newTreeType(name, color, texture)
		t.treeTypes = append(t.treeTypes, tree)
	}
	return tree
}

func (t *treeFactory) find(name, color, texture string) *treeType {
	for _, tree := range t.treeTypes {
		if tree.name == name && tree.color == color && tree.texture == texture {
			return tree
		}
	}

	return nil
}
