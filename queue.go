package ehtml

import "golang.org/x/net/html"

type (
	queue struct {
		start, end *node
		length     int
	}
	node struct {
		value *html.Node
		next  *node
	}
)

// Create a new queue
func newNodeQueue() *queue {
	return &queue{nil, nil, 0}
}

// Take the next item off the front of the queue
func (this *queue) Dequeue() *html.Node {
	if this.length == 0 {
		return nil
	}
	n := this.start
	if this.length == 1 {
		this.start = nil
		this.end = nil
	} else {
		this.start = this.start.next
	}
	this.length--
	return n.value
}

// Put an item on the end of a queue
func (this *queue) Enqueue(value *html.Node) {
	n := &node{value, nil}
	if this.length == 0 {
		this.start = n
		this.end = n
	} else {
		this.end.next = n
		this.end = n
	}
	this.length++
}

// Return the number of items in the queue
func (this *queue) Len() int {
	return this.length
}

// Return the first item in the queue without removing it
func (this *queue) Peek() *html.Node {
	if this.length == 0 {
		return nil
	}
	return this.start.value
}
