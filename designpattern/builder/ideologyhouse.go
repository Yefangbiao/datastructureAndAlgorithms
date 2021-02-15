package main

type ideologyHose struct {
	house
}

func NewIdeologyHose() *ideologyHose {
	return &ideologyHose{}
}

func (h *ideologyHose) setWall() {
	h.wall = "good wall"
}

func (h *ideologyHose) setWindow() {
	h.window = "good window"
}

func (h *ideologyHose) setDoor() {
	h.door = "good door"
}

func (h *ideologyHose) getResult() house {
	return house{
		wall:   h.wall,
		window: h.window,
		door:   h.door,
	}
}
