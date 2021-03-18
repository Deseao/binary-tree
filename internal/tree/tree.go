package tree

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Value int //TODO: make this interface{}?
	Right *Tree
}

func New(n int) *Tree { //TODO: make this accept interface{}
	return &Tree{Value: n}
}

func (t *Tree) Insert(n int) *Tree {
	if t == nil {
		return New(n)
	}
	if n <= t.Value {
		t.Left = t.Left.Insert(n)
	} else {
		t.Right = t.Right.Insert(n)
	}
	return t
}

func (t *Tree) Print() string {
	if t == nil {
		return "()"
	}
	result := ""
	if t.Left != nil {
		result += t.Left.Print() + " "
	}
	result += fmt.Sprint(t.Value)
	if t.Right != nil {
		result += t.Right.Print() + " "
	}
	return "(" + result + ")"
}
