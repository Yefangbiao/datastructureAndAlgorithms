package main

func main() {
	forest := NewForest()
	forest.PlantTree(3, 5, "greenTree", "green", "nice")
	forest.PlantTree(43, 55, "greenTree", "green", "nice")
	forest.PlantTree(35, 5, "yellowTree", "yellow", "bad")
	forest.PlantTree(6, 132, "yellowTree", "yellow", "bad")
	forest.PlantTree(6236, 6578, "yellowTree", "yellow", "excellent")
	forest.Draw()
}
