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
func (q *queue) Dequeue() *html.Node {
	if q.length == 0 {
		return nil
	}
	n := q.start
	if q.length == 1 {
		q.start = nil
		q.end = nil
	} else {
		q.start = q.start.next
	}
	q.length--
	return n.value
}

// Put an item on the end of a queue
func (q *queue) Enqueue(value *html.Node) {
	n := &node{value, nil}
	if q.length == 0 {
		q.start = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}
	q.length++
}

// EnqueueSlice takes a whole slice of nodes and enqueues them.
func (q *queue) EnqueueSlice(values []*html.Node) {
	for _, v := range values {
		q.Enqueue(v)
	}
}

// Return the number of items in the queue
func (q *queue) Len() int {
	return q.length
}

// Return the first item in the queue without removing it
func (q *queue) Peek() *html.Node {
	if q.length == 0 {
		return nil
	}
	return q.start.value
}

// Slice returns the queue as a slice of nodes.
func (q *queue) Slice() []*html.Node {
	s := make([]*html.Node, q.Len())

	for i := 0; i < q.Len(); i++ {
		s[i] = q.Dequeue()
	}

	return s
}
