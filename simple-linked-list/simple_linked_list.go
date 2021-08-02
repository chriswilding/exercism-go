package linkedlist

import (
	"errors"
)

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(numbers []int) *List {
	l := List{}

	if numbers == nil || len(numbers) == 0 {
		return &l
	}

	for _, n := range numbers {
		l.Push(n)
	}

	return &l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(n int) {
	e := Element{
		data: n,
		next: l.head,
	}
	l.size++
	l.head = &e
}

func (l *List) Pop() (int, error) {
	if l.head == nil || l.size == 0 {
		return 0, errors.New("list is empty")
	}
	head := l.head
	next := head.next
	l.head = next
	l.size--
	return head.data, nil
}

func (l *List) Array() []int {
	if l.head == nil || l.size == 0 {
		return nil
	}

	a := make([]int, l.size)
	var head = l.head
	i := l.size - 1

	for i >= 0 {
		a[i] = head.data
		i -= 1
		head = head.next
	}
	return a
}

func (l *List) Reverse() *List {
	var next *Element
	var previous *Element
	current := l.head

	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}

	l.head = previous

	return l
}
