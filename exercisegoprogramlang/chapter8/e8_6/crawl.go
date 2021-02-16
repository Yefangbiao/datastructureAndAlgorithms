package main

import (
	"flag"
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)

type depth int

func (d *depth) String() string {
	return fmt.Sprintln(d)
}

func (d *depth) Set(s string) error {
	var depth int
	_, err := fmt.Sscanf(s, "%d", &depth)
	return err
}

func DepthFlag(name string, value depth, usage string) *depth {
	d := &value
	flag.CommandLine.Var(d, name, usage)
	return d
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

var d = DepthFlag("depth", 3, "depth")

type linkList struct {
	url []string
	d   int
}

func main() {
	flag.Parse()

	worklist := make(chan linkList)
	var n int
	n++
	go func() {
		urls := os.Args[1:]
		worklist <- linkList{urls, 1}
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list.url {
			if list.d > int(*d) {
				break
			}
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					urls := crawl(link)
					worklist <- linkList{urls, list.d}
				}(link)
			}
		}
	}
}
