package main

import "fmt"

type Node struct {
	value int
}

type Stack struct {
	nodes []*Node // slice of pointers to Nodes
	size  int
}

func (n *Node) ToString() string {
	return fmt.Sprint(n.value)
}

func newStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(ele int) {
	n := new(Node)
	n.value = ele
	s.nodes = append(s.nodes[:s.size], n)
	s.size += 1
}

func (s *Stack) pop() *Node {
	if s.size == 0 {
		return nil
	}
	s.size -= 1
	n := s.nodes[s.size]
	s.nodes = s.nodes[:s.size]
	return n
}

func main() {
	stack := newStack()
	fmt.Println(stack)
	for i, e := range stack.nodes {
		fmt.Println(i, e.ToString())
	}
	stack.push(10)
	stack.push(25)
	stack.push(3)
	stack.push(43)
	fmt.Println(stack)
	for i, e := range stack.nodes {
		fmt.Println(i, e.ToString())
	}
	fmt.Println(stack.pop(), stack.pop(), stack.pop())
	fmt.Println(stack)
	for i, e := range stack.nodes {
		fmt.Println(i, e.ToString())
	}
	fmt.Println(stack.pop(), stack.pop())
}
