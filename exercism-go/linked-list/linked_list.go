package linkedlist

import "errors"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type List struct {
	len        int
	head, tail *Node
}
type Node struct {
	Value      interface{}
	next, prev *Node
}

func NewList(elements ...interface{}) *List {
	var l List = List{}
	var prev *Node
	for _, v := range elements {
		// create new node with current value and prev pointer
		currNode := Node{Value: v, prev: prev}

		// if first node, assign head
		if l.len == 0 {
			l.head = &currNode
			// else update the tail's next pointer
		} else {
			l.tail.next = &currNode
		}

		l.tail = &currNode
		prev = &currNode
		l.len++
	}
	return &l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	newNode := &Node{Value: v}
	if l.len == 0 {
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
	}
	l.head = newNode
	l.len++
}

func (l *List) Push(v interface{}) {
	newNode := &Node{Value: v}
	if l.len == 0 {
		l.head = newNode
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
	}
	l.tail = newNode
	l.len++
}

func (l *List) Shift() (interface{}, error) {
	if l.len == 0 {
		return nil, errors.New("List is empty")
	}
	currNode := l.head
	l.len--
	if l.len == 0 {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}
	return currNode.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.len == 0 {
		return nil, errors.New("List is empty")
	}
	currNode := l.tail
	l.len--
	if l.len == 0 {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}
	return currNode.Value, nil
}

func (l *List) Reverse() {
	for p := l.head; p != nil; {
		p.next, p.prev = p.prev, p.next
		p = p.prev // NOTE: instead of p = p.next because we swapped, this hurt my brain so bad :)
	}
	l.tail, l.head = l.head, l.tail
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
