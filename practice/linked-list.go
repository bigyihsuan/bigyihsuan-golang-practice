package main

import "fmt"

type Node[E any] struct {
	value E
	prev  *Node[E]
	next  *Node[E]
}

type LinkedList[E any] struct {
	head *Node[E]
	tail *Node[E]
	size uint
}

func newNode[E any](value E) *Node[E] {
	node := new(Node[E])
	node.value = value
	return node
}

func newLinkedList[E any]() LinkedList[E] {
	return LinkedList[E]{head: nil, tail: nil, size: 0}
}

func (list *LinkedList[E]) append(value E) {
	// cases
	// - list is empty
	// - list has elements

	node := newNode(value)
	// list is empty
	if list.head == nil && list.tail == nil {
		list.head = node
		list.tail = node
		list.size++
	} else {
		// list has elements
		node.prev, list.tail.next = list.tail, node
		list.tail = list.tail.next
	}
	list.size++
}

func (list *LinkedList[E]) pop() *E {

	// cases
	// - list is empty
	// - list has 1 element
	// - list has n elements

	node := list.tail
	// empty list
	if list.size == 0 && list.head == nil && list.tail == nil {
		return nil
	} else if list.size == 1 {
		list.tail = nil
		list.head = nil
	} else {
		list.tail = node.prev
		list.tail.next = nil
		node.prev = nil
	}
	list.size--
	return &node.value
}

func main() {
	arr := []int{1, 2, 3, 5, 6, 7}

	list := newLinkedList[int]()

	for _, e := range arr {
		list.append(e)
	}

	for curr := list.head; curr != nil; curr = curr.next {
		fmt.Println(curr.value)
	}
	fmt.Println(*list.pop(), *list.pop(), *list.pop())
	list.pop()
	list.pop()
	list.pop()
	fmt.Println(*list.pop(), *list.pop(), *list.pop())
}
