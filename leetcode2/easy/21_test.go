package easy

import "fmt"

func Example21() {

	// list1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	// list2 := ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}

	list1 := ListNode{Val: 2, Next: nil}
	list2 := ListNode{Val: 1, Next: nil}
	list := mergeTwoLists(&list1, &list2)

	next := list
	for next != nil {
		fmt.Println(next.Val)
		next = next.Next
	}
	// Output:
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	next1, next2 := l1, l2
	l := new(ListNode)
	next := l
	for next1 != nil || next2 != nil {
		if next1 != nil && (next2 == nil || next1.Val <= next2.Val) {
			next.Next = &ListNode{Val: next1.Val, Next: nil}
			next = next.Next
			next1 = next1.Next
		}
		if next2 != nil && (next1 == nil || next2.Val <= next1.Val) {
			next.Next = &ListNode{Val: next2.Val, Next: nil}
			next = next.Next
			next2 = next2.Next
		}
	}
	return l.Next
}

// TODO:
// 试试递归
