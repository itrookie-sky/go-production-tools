package reorderList

import "testing"

func TestReorderList(t *testing.T) {
	var a, b, c, d ListNode
	a.Val = 1
	a.Next = &b
	b.Val = 2
	b.Next = &c
	c.Val = 3
	c.Next = &d
	d.Val = 4
	// d.Next = nil

	ReorderList(&a)
	a.Print()
}
