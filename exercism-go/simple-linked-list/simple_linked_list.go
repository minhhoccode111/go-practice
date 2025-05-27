package linkedlist

import "errors"

// Define the List and Element types here.
type List struct {
	size int
	head *Element
	tail *Element
}
type Element struct {
	val  int
	next *Element
	prev *Element
}

func New(elements []int) *List {
	result := List{}
	for _, v := range elements {
		result.Push(v)
	}
	return &result
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	curr := Element{val: element}
	if l.size == 0 {
		l.head = &curr
		l.tail = &curr
		l.size++
		return
	}
	l.tail.next = &curr
	curr.prev = l.tail
	l.tail = &curr
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size == 0 {
		return 0, errors.New("can't pop empty list")
	}
	curr := l.tail
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	}
	l.size--
	return curr.val, nil
}

func (l *List) Array() []int {
	result := []int{}
	for p := l.head; p != nil; p = p.next {
		result = append(result, p.val)
	}
	return result
}

func (l *List) Reverse() *List {
	for p := l.head; p != nil; p = p.prev {
		p.next, p.prev = p.prev, p.next
	}
	l.head, l.tail = l.tail, l.head
	return l
}
