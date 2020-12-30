package reorderList

import "fmt"

//链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

//输出日志
func (node *ListNode) Print() {
	fmt.Print(node.Val, "->")
	if node.Next != nil {
		node.Next.Print()
	} else {
		fmt.Println("")
	}
}

/*
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

type ListNode struct {
	Val int
	Next *ListNode
}
*/
func ReorderList(head *ListNode) {
	var last, second, secondLast *ListNode

	current := head
	i := 0

	for current.Next != nil && current.Next.Next != nil {
		if i == 1 {
			second = current
		}
		current = current.Next
		i++
	}
	secondLast = current
	last = current.Next

	head.Next = last
	last.Next = second
	secondLast.Next = nil
}

//官方题解
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}

//力扣官方答案
func LeetcodeReorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
}
