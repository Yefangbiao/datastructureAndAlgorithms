package e5_8

import (
	"fmt"
	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var elem *html.Node
	forEachNode(doc,
		func(n *html.Node) bool {
			if n.Type == html.ElementNode {
				for _, a := range n.Attr {
					if a.Key == "id" && a.Val == id {
						elem = n
						return false
					}
				}
			}
			return true
		},
		nil)
	return elem
}
