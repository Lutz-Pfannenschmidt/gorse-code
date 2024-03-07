package main

import "fmt"

type Tree struct {
	Left  *Tree
	Right *Tree
	Value rune
}

func (t *Tree) Insert(path string, value rune) {
	if len(path) == 0 {
		t.Value = value
		return
	}

	if path[0] == '.' {
		if t.Left == nil {
			t.Left = &Tree{}
		}
		t.Left.Insert(path[1:], value)
	} else {
		if t.Right == nil {
			t.Right = &Tree{}
		}
		t.Right.Insert(path[1:], value)
	}
}

func (t *Tree) Get(path string) (rune, error) {
	if len(path) == 0 {
		return t.Value, nil
	}

	if path[0] == '.' {
		if t.Left == nil {
			return 0, fmt.Errorf("not found")
		}
		return t.Left.Get(path[1:])
	} else {
		if t.Right == nil {
			return 0, fmt.Errorf("not found")
		}
		return t.Right.Get(path[1:])
	}
}
