package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	doc := fetch("https://golang.org")

	counts := make(map[string]int)
	visit(counts, doc)
	for k, v := range counts {
		fmt.Printf("%s = %d\n", k, v)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(counts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(counts, c)
	}
}

func fetch(url string) *html.Node {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	return doc
}
