package main

type director struct {
	iBuild
}

func NewDirector() *director {
	return &director{}
}

func (d *director) SetBuilder(b iBuild) {
	d.iBuild = b
}

func (d *director) MakeHouse() house {
	d.setWall()
	d.setWindow()
	d.setDoor()
	return d.getResult()
}
