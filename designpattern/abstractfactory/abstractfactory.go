package main

import "fmt"

type abstractFactory interface {
	newShirt(size int) shirt
	newShoe(size int) shoe
}

func GetFactory(brand string) (abstractFactory, error) {
	switch brand {
	case "nike":
		return nike{}, nil
	case "adidas":
		return adidas{}, nil
	}
	return nil, fmt.Errorf("GetFactory: ")
}
