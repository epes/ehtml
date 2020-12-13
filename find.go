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

// FindNested applies multiple testing functions in sequence to find
// the nodes which satisfy all of them.
func FindNested(root *html.Node, fns []FindTestFn) []*html.Node {
	q := newNodeQueue()
	q.Enqueue(root)

	for _, fn := range fns {
		newQ := newNodeQueue()

		for node := q.Dequeue(); node != nil; node = q.Dequeue() {
			newQ.EnqueueSlice(Find(node, fn))
		}

		q = newQ
	}

	return q.Slice()
}

func getClassTestingFn(class string) FindTestFn {
	class = strings.ToLower(class)

	return func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "class" && strings.ToLower(a.Val) == class {
					return true
				}
			}
		}

		return false
	}
}

// FindClass is sugar for Find for a specific class attribute.
func FindClass(root *html.Node, class string) []*html.Node {
	class = strings.ToLower(class)

	return Find(root, getClassTestingFn(class))
}

// FindClasses finds the elements which match the specified nested classes.
func FindClasses(root *html.Node, classes []string) []*html.Node {
	var fns []FindTestFn

	for _, c := range classes {
		fns = append(fns, getClassTestingFn(c))
	}

	return FindNested(root, fns)
}

func getTagTestingFn(tag string) FindTestFn {
	tag = strings.ToLower(tag)

	return func(n *html.Node) bool {
		fmt.Println(n.Data)

		if n.Type == html.ElementNode && strings.ToLower(n.Data) == tag {
			return true
		}

		return false
	}
}

// FindTag is sugar for Find for a specific element name.
func FindTag(root *html.Node, tag string) []*html.Node {

	return Find(root, getTagTestingFn(tag))
}

// FindTags finds the elements which match the specified nested tags.
func FindTags(root *html.Node, tags []string) []*html.Node {
	var fns []FindTestFn

	for _, t := range tags {
		fns = append(fns, getTagTestingFn(t))
	}

	return FindNested(root, fns)
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
