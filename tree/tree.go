package tree

import (
	"fmt"
	"strconv"
)

// Узел
type Node struct {
	key   int
	left  *Node
	right *Node
}

// Дерево
type Tree struct {
	root *Node
}

// Создание дерева
func New() *Tree {
	return &Tree{
		root: nil,
	}
}

// Вставка элемента
func (t *Tree) Insert(v int) {
	newNode := &Node{
		key: v,
	}

	if t.root == nil {
		t.root = newNode
		return
	}

	currentNode := t.root
	parentNode := t.root

	for {
		parentNode = currentNode
		if v < currentNode.key {
			currentNode = currentNode.left
			if currentNode == nil {
				parentNode.left = newNode
				return
			}
		} else if v >= currentNode.key {
			currentNode = currentNode.right
			if currentNode == nil {
				parentNode.right = newNode
				return
			}
		}
	}
}

func (t *Tree) Find(v int) bool {
	return t.root.findElement(v)
}

func (r *Node) findElement(v int) bool {
	if r != nil {
		if v == r.key {
			return true
		} else if v > r.key {
			return r.right.findElement(v)
		} else {
			return r.left.findElement(v)
		}
	}
	return false
}

func (t *Tree) Delete(v int) bool {
	if !t.Find(v) || t.root == nil {
		return false
	}

	if t.root.key == v {
		tempRoot := &Node{0, nil, nil}
		tempRoot.left = t.root
		r := del(t.root, tempRoot, v)
		t.root = tempRoot.left
		return r
	}
	return del(t.root.left, t.root, v) || del(t.root.right, t.root, v)
}

func minValue(root *Node) int {
	if root != nil {
		if root.left == nil {
			return root.key
		}
		return minValue(root.left)
	}
	return root.key
}

func del(root *Node, parent *Node, value int) bool {
	switch {
	case root.key == value:
		if root.left != nil && root.right != nil {
			root.key = minValue(root.right)
			return del(root.right, root, root.key)
		}
		link(parent, root)
		return true
	case root.key > value:
		if root.left == nil {
			return false
		}
		return del(root.left, root, value)
	case root.key < value:
		if root.right == nil {
			return false
		}
		return del(root.right, root, value)
	}
	return false
}

func link(parent *Node, root *Node) {
	if parent.left == root {
		if root.left != nil {
			parent.left = root.left
		} else {
			parent.left = root.right
		}
	} else if parent.right == root {
		if root.left != nil {
			parent.right = root.left
		} else {
			parent.right = root.right
		}
	}
}

func (t *Tree) Print() {
	fmt.Println(t.root.printNode() + "\n")
}

func (r *Node) printNode() string {
	result := ""
	if r != nil {
		result += r.left.printNode() + " " + strconv.Itoa(r.key) + " " + r.right.printNode()
	}
	return result
}
