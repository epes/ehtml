package ehtml

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// FindTestFn is a testing function to be applied
// to an *html.Node.
type FindTestFn func(*html.Node) bool

// Find applies the tester function to all nodes within the specified
// root subtree and returns all nodes that pass.
func Find(root *html.Node, fn FindTestFn) []*html.Node {
	q := newNodeQueue()
	var res []*html.Node

	for node := root; node != nil; node = q.Dequeue() {
		if fn(node) {
			res = append(res, node)
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			q.Enqueue(c)
		}
	}

	return res
}

// FindClass is sugar for Find for a specific class attribute.
func FindClass(root *html.Node, class string) []*html.Node {
	class = strings.ToLower(class)

	return Find(root, func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "class" && strings.ToLower(a.Val) == class {
					return true
				}
			}
		}

		return false
	})
}

// FindTags is sugar for Find for a specific element name.
func FindTags(root *html.Node, tag string) []*html.Node {
	tag = strings.ToLower(tag)

	return Find(root, func(n *html.Node) bool {
		fmt.Println(n.Data)

		if n.Type == html.ElementNode && strings.ToLower(n.Data) == tag {
			return true
		}

		return false
	})
}

// FindText finds all html.TextNode within the specified root tree.
func FindText(root *html.Node) []*html.Node {
	return Find(root, func(n *html.Node) bool {
		if n.Type == html.TextNode {
			return true
		}

		return false
	})
}
