package main

type visitor interface {
	visitCircle(*circle)
	visitorSquare(*square)
}
