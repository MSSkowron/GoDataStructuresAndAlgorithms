package singlylinkedlist

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type node[T constraints.Ordered] struct {
	value T
	next  *node[T]
}

func newNode[T constraints.Ordered](value T) *node[T] {
	return &node[T]{
		value: value,
	}
}

type LinkedList[T constraints.Ordered] struct {
	head   *node[T]
	length int
}

func New[T constraints.Ordered]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Prepend(value T) {
	ll.length++

	n := newNode[T](value)
	n.next = ll.head
	ll.head = n
}

func (ll *LinkedList[T]) Append(value T) {
	ll.length++

	if ll.head == nil {
		ll.head = newNode[T](value)
		return
	}

	curr := ll.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode[T](value)
}

func (ll *LinkedList[T]) Delete(value T) {
	if ll.head == nil {
		return
	}

	if ll.head.value == value {
		ll.head = ll.head.next
		ll.length--
		return
	}

	curr := ll.head
	for curr.next != nil {
		if curr.next.value == value {
			curr.next = curr.next.next
			ll.length--
			return
		}

		curr = curr.next
	}
}

func (ll *LinkedList[T]) Print() {
	if ll.head == nil {
		fmt.Println("nil")
		return
	}

	curr := ll.head
	for curr != nil {
		fmt.Printf("%v -> ", curr.value)
		curr = curr.next
	}
	fmt.Printf("nil")
}
