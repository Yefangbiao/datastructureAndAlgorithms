package main

import "fmt"

type nike struct {
}

// newNikeShirt return a nike shirt
// size: the size of shirt
func (n nike) newShirt(size int) shirt {
	return nikeShirt{iShirt: iShirt{
		producer: "nike",
		size:     size,
	}}
}

// newNikeShore return a nike shore
// size: the size of shoe
func (n nike) newShoe(size int) shoe {
	return &nikeShoe{iShoe: iShoe{
		producer: "nike",
		size:     size,
	}}
}

// + nikeShore
type nikeShoe struct {
	iShoe
}

func (n nikeShoe) Action() {
	fmt.Println("穿上了nike鞋")
}

// - nikeShore

// + nikeShirt
type nikeShirt struct {
	iShirt
}

func (n nikeShirt) Action() {
	fmt.Println("穿上了nike衣服")
}

// - nikeShirt
