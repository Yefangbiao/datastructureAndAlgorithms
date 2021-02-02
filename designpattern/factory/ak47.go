package main

type ak47 struct {
	gun
}

func newAk47() *ak47 {
	return &ak47{gun: gun{
		name:  "ak47",
		power: 41,
	}}
}
