package main

type normalHouse struct {
	house
}

func NewNormalHouse() *normalHouse {
	return &normalHouse{}
}

func (h *normalHouse) setWall() {
	h.wall = "normal wall"
}

func (h *normalHouse) setWindow() {
	h.window = "normal window"
}

func (h *normalHouse) setDoor() {
	h.door = "normal door"
}

func (h *normalHouse) getResult() house {
	return house{
		wall:   h.wall,
		window: h.window,
		door:   h.door,
	}
}
