package main

import "fmt"

type Node struct {
	value       int
	left, right *Node
}

func NewNode(value int) *Node {
	node := new(Node)
	node.value = value
	return node
}

// for fmt.Stringer
func (n Node) String() string {
	var l, r string
	if n.left == nil {
		l = ""
	} else {
		l = fmt.Sprint(n.left)
	}
	if n.right == nil {
		r = ""
	} else {
		r = fmt.Sprint(n.right)
	}
	if n.left == nil && n.right == nil {
		return fmt.Sprint(n.value)
	}
	return fmt.Sprintf("%v(%v,%v)", n.value, l, r)
}

// returned by IsBinarySearchTree, stores which nodes causes the offending node to not be a BST
type NotBSTError struct {
	offendingNode, badLeftChild, badRightChild *Node
}

// for error.Error interface
func (err NotBSTError) Error() string {
	outstr := fmt.Sprintf("NotBSTError: tree with key %v is not a BST because ", err.offendingNode.value)
	switch {
	case err.badLeftChild != nil:
		outstr += fmt.Sprintf("left node %v >= root", err.badLeftChild.value)
	case err.badRightChild != nil:
		outstr += fmt.Sprintf("right node %v <= root", err.badRightChild.value)
	default:
		outstr += fmt.Sprintf("left node %v >= root, and ", err.badLeftChild.value)
		outstr += fmt.Sprintf("right node %v <= root", err.badRightChild.value)
	}
	return outstr
}

// checks if a given node and its children is a BST
// BSTs are defined such that left < root < right
func (root Node) IsBinarySearchTree() (bool, *NotBSTError) {
	switch {
	// both children are nil
	case root.left == nil && root.right == nil:
		return true, nil
		// left is not nil, right is nil
	case root.left != nil && root.right == nil:
		if root.left.value >= root.value {
			return false, &NotBSTError{&root, root.left, nil}
		}
		return root.left.IsBinarySearchTree()
		// left is nil, right is not nil
	case root.right != nil && root.left == nil:
		if root.right.value <= root.value {
			return false, &NotBSTError{&root, root.right, nil}
		}
		return root.right.IsBinarySearchTree()
	// both are not nil
	default:
		switch {
		case root.left.value >= root.value && root.right.value <= root.value:
			return false, &NotBSTError{&root, root.left, root.right}
		case root.left.value >= root.value:
			return false, &NotBSTError{&root, root.left, nil}
		case root.right.value <= root.value:
			return false, &NotBSTError{&root, nil, root.right}
		}
		l, errl := root.left.IsBinarySearchTree()
		if errl != nil {
			return false, errl
		}
		r, errr := root.right.IsBinarySearchTree()
		if errr != nil {
			return false, errr
		}
		return l && r, nil
	}
}

func (root Node) IsValueInTree(value int) bool {
	switch {
	case root.value == value:
		// found value
		return true
	case root.left == nil && root.right == nil:
		// children are null
		return false
	case root.left != nil && root.right == nil:
		// left is not nil
		return root.left.IsValueInTree(value)
	case root.right != nil && root.left == nil:
		// right is not nil
		return root.right.IsValueInTree(value)
	default:
		// otherwise
		return root.left.IsValueInTree(value) || root.right.IsValueInTree(value)
	}
}

func (root Node) IsValueInBST(value int) bool {
	switch {
	case root.value == value:
		// found value
		return true
	case root.value > value && root.left != nil:
		// value is less/left
		return root.left.IsValueInBST(value)
	case root.value < value && root.right != nil:
		// value is more/right
		return root.right.IsValueInBST(value)
	default:
		// no children and not equal
		return false
	}
}

func main() {
	var notBST = NewNode(2)
	notBST.left = NewNode(7)
	notBST.left.left = NewNode(2)
	notBST.left.right = NewNode(6)
	notBST.left.right.left = NewNode(5)
	notBST.left.right.right = NewNode(11)
	notBST.right = NewNode(5)
	notBST.right.right = NewNode(9)
	notBST.right.right.left = NewNode(4)
	fmt.Println(notBST)
	ok, err := notBST.IsBinarySearchTree()
	fmt.Println(ok, err)

	bst := NewNode(8)
	bst.left = NewNode(3)
	bst.left.left = NewNode(1)
	bst.left.right = NewNode(6)
	bst.left.right.left = NewNode(4)
	bst.left.right.right = NewNode(7)
	bst.right = NewNode(10)
	bst.right.right = NewNode(14)
	bst.right.right.left = NewNode(13)
	fmt.Println(bst)
	ok, err = bst.IsBinarySearchTree()
	fmt.Println(ok, err)

	fmt.Printf("notBST.IsValueInTree(10): %v\n", notBST.IsValueInTree(10))
	fmt.Printf("notBST.IsValueInTree(11): %v\n", notBST.IsValueInTree(11))

	fmt.Printf("bst.IsValueInTree(2): %v\n", bst.IsValueInTree(2))
	fmt.Printf("bst.IsValueInTree(10): %v\n", bst.IsValueInTree(10))

	fmt.Printf("bst.IsValueInBST(2): %v\n", bst.IsValueInBST(2))
	fmt.Printf("bst.IsValueInBST(10): %v\n", bst.IsValueInBST(10))
}
