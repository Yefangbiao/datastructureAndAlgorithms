package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	doc := fetch("https://golang.org")

	for _, link := range visit(nil, doc) {
		fmt.Printf("%s\n", link)
	}
}

var filter = map[string]string{
	"a":      "href",
	"img":    "src",
	"script": "src",
}

func visit(links []string, n *html.Node) []string {
	for k, v := range filter {
		if n.Type == html.ElementNode && n.Data == k {
			for _, a := range n.Attr {
				if a.Key == v {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

// 这里我把fetch封装成一个函数调用
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
