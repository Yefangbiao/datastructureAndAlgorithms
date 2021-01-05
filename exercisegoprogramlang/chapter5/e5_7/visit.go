package e5_7

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func FindLinks(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("findLinks: %v", err)
	}
	// 随后关闭
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s: %v", url, resp.StatusCode)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("parsing %s as HTML: %v", url, err)
	}
	visit(doc, startElement, endElement)
	return nil
}

func visit(n *html.Node, pre, post func(str string)) {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if pre != nil {
				pre("<" + a.Key + ">")
			}
			fmt.Print(a.Val)
			if post != nil {
				post("<" + a.Key + ">")
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, pre, post)
	}

}

func startElement(str string) {
	fmt.Print(str)
}

func endElement(str string) {
	fmt.Print(str + "\n")
}
