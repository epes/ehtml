package ehtml

import (
	"strings"

	"golang.org/x/net/html"
)

// FindTestFn is a testing function to be applied
// to an *html.Node.
type FindTestFn func(*html.Node) bool

// Find applies the tester function to all nodes within the specified
// root subtree and returns all nodes that pass.
func Find(root *html.Node, fn FindTestFn) []*html.Node {
	if root == nil {
		return nil
	}

	q := newNodeQueue()
	var res []*html.Node

	for c := root.FirstChild; c != nil; c = c.NextSibling {
		q.Enqueue(c)
	}

	for node := q.Dequeue(); node != nil; node = q.Dequeue() {
		if fn(node) {
			res = append(res, node)
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			q.Enqueue(c)
		}
	}

	return res
}

func findNested(root *html.Node, fns []FindTestFn) []*html.Node {
	q := newNodeQueue()

	if root != nil {
		q.Enqueue(root)
	}

	for _, fn := range fns {
		newQ := newNodeQueue()

		for node := q.Dequeue(); node != nil; node = q.Dequeue() {
			nested := Find(node, fn)

			if nested != nil {
				newQ.EnqueueSlice(nested)
			}
		}

		q = newQ
	}

	return q.Slice()
}

// FindJ finds nested nodes given selectors similar to JavaScript querySelectors.
func FindJ(root *html.Node, joinedSelectors string) []*html.Node {
	if root == nil || joinedSelectors == "" {
		return nil
	}

	var fns []FindTestFn

	selectors := strings.Split(joinedSelectors, " ")

	for _, s := range selectors {
		if len(s) == 0 {
			continue
		}

		switch s[0] {
		case '.':
			fns = append(fns, getClassTestingFn(s[1:]))
		case '*':
			fns = append(fns, textTestingFn)
		default:
			fns = append(fns, getTagTestingFn(s))
		}
	}

	return findNested(root, fns)
}

func getClassTestingFn(class string) FindTestFn {
	class = strings.ToLower(class)

	return func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "class" {
					classes := strings.Split(a.Val, " ")
					for _, c := range classes {
						if strings.ToLower(c) == class {
							return true
						}
					}
				}
			}
		}

		return false
	}
}

func getTagTestingFn(tag string) FindTestFn {
	tag = strings.ToLower(tag)

	return func(n *html.Node) bool {
		return n.Type == html.ElementNode && strings.ToLower(n.Data) == tag
	}
}

func textTestingFn(n *html.Node) bool {
	return n.Type == html.TextNode
}
