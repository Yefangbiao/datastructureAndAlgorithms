package main

import "fmt"

func main() {
	roundhole := roundHole{radius: 20}
	rp := roundPeg{radius: 23}
	fmt.Println(roundhole.fits(&rp))

	sp := squarePeg{width: 4}
	sAdapter := squareAdapter{&sp}
	fmt.Println(roundhole.fits(&sAdapter))
}
